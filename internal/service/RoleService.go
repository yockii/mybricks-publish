package service

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	logger "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"manatee-publish/internal/constant"
	"manatee-publish/internal/model"
	"manatee-publish/pkg/cache"
	"manatee-publish/pkg/common"
	"manatee-publish/pkg/database"
	"manatee-publish/pkg/util"
)

var RoleService = newRoleService()

type roleService struct {
	common.BaseService[*model.Role]
}

func newRoleService() *roleService {
	s := new(roleService)
	s.BaseService = common.BaseService[*model.Role]{
		Service: s,
	}
	return s
}

func (*roleService) Model() *model.Role {
	return new(model.Role)
}

// Update 更新角色基本信息
func (s *roleService) Update(instance *model.Role, tx ...*gorm.DB) (count int64, err error) {
	count, err = s.BaseService.Update(instance, tx...)
	if count > 0 {
		s.removeCache(instance.ID)
	}
	return
}

// Delete 删除角色
func (s *roleService) Delete(instance *model.Role, tx ...*gorm.DB) (count int64, err error) {
	count, err = s.BaseService.Delete(instance, tx...)
	if count > 0 {
		s.removeCache(instance.ID)
	}
	return
}

// ResourceCodes 获取角色的资源编码列表
func (s *roleService) ResourceCodes(roleId uint64) (list []string, err error) {
	list = make([]string, 0)
	err = database.DB.Model(&model.RoleResource{}).Where(&model.RoleResource{RoleID: roleId}).Pluck("resource_code", &list).Error
	if err != nil {
		logger.Errorln(err)
		return
	}
	return
}

func (s *roleService) DispatchResources(roleID uint64, ResourceCodeList []string) (success bool, err error) {
	if roleID == 0 {
		err = errors.New("roleID is required")
		return
	}
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		// 删除旧的
		err = tx.Where(&model.RoleResource{RoleID: roleID}).Delete(&model.RoleResource{}).Error
		if err != nil {
			logger.Errorln(err)
			return err
		}
		// 添加新的
		for _, resourceCode := range ResourceCodeList {
			roleResource := &model.RoleResource{
				BaseModel:    common.BaseModel{ID: util.SnowflakeId()},
				RoleID:       roleID,
				ResourceCode: resourceCode,
			}
			err = tx.Create(roleResource).Error
			if err != nil {
				logger.Errorln(err)
				return err
			}
		}
		return nil
	})
	if err != nil {
		logger.Errorln(err)
		return
	}
	success = true
	return
}

// SetDefault 设置默认角色
func (*roleService) SetDefault(id uint64) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Role{}).Where("default_role=?", 1).Updates(&model.Role{DefaultRole: -1}).Error; err != nil {
			logger.Errorln(err)
			return err
		}
		if err := tx.Model(&model.Role{BaseModel: common.BaseModel{ID: id}}).Updates(&model.Role{DefaultRole: 1}).Error; err != nil {
			logger.Errorln(err)
			return err
		}
		return nil
	})
}

func (s *roleService) removeCache(id uint64) {
	conn := cache.Get()
	defer func(conn redis.Conn) {
		_ = conn.Close()
	}(conn)
	_, _ = conn.Do("HDEL", constant.RedisKeyRoleDataPerm, id)
}
