package service

import (
	"errors"
	logger "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"manatee-publish/internal/domain"
	"manatee-publish/internal/model"
	"manatee-publish/pkg/common"
	"manatee-publish/pkg/database"
	"manatee-publish/pkg/util"
)

var UserService = newUserService()

type userService struct {
	common.BaseService[*model.User]
}

func newUserService() *userService {
	s := new(userService)
	s.BaseService = common.BaseService[*model.User]{
		Service: s,
	}
	return s
}

func (*userService) Model() *model.User {
	return new(model.User)
}

// LoginWithUsernameAndPassword 用户登录
func (s *userService) LoginWithUsernameAndPassword(username, password string) (instance *model.User, passwordNotMatch bool, err error) {
	instance = new(model.User)
	err = database.DB.Where(&model.User{Username: username}).First(instance).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			passwordNotMatch = true
			return
		}
		logger.Errorln(err)
		return
	}
	if instance.Status != model.UserStatusNormal {
		err = errors.New("用户已被禁用")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(instance.Password), []byte(password))
	if err != nil {
		passwordNotMatch = true
		err = nil
		return
	}
	// 完成后密码置空
	instance.Password = ""
	return
}

// AddDomain 添加用户
func (s *userService) AddDomain(instance *domain.UserDomain, tx ...*gorm.DB) (duplicated bool, err error) {
	if instance.Username == "" {
		err = errors.New("username is required")
		return
	}
	var c int64
	err = database.DB.Model(&model.User{}).Where(&model.User{Username: instance.Username}).Count(&c).Error
	if err != nil {
		logger.Errorln(err)
		return
	}
	if c > 0 {
		duplicated = true
		return
	}

	instance.ID = util.SnowflakeId()
	if instance.Password != "" {
		pwd, _ := bcrypt.GenerateFromPassword([]byte(instance.Password), bcrypt.DefaultCost)
		instance.Password = string(pwd)
	}
	instance.Status = model.UserStatusNormal

	// 获取默认角色
	defaultRole := &model.Role{DefaultRole: 1}
	if err = database.DB.Where(defaultRole).First(defaultRole).Error; err != nil {
		logger.Errorln(err)
		return
	}
	if defaultRole != nil && defaultRole.ID > 0 {
		if len(tx) == 0 {
			// 添加用户的同时要添加默认角色
			err = database.DB.Transaction(func(tx *gorm.DB) error {
				if err = tx.Create(instance.User).Error; err != nil {
					logger.Errorln(err)
					return err
				}
				userRole := &model.UserRole{
					BaseModel: common.BaseModel{ID: util.SnowflakeId()},
					UserID:    instance.ID,
					RoleID:    defaultRole.ID,
				}
				if err = tx.Create(userRole).Error; err != nil {
					logger.Errorln(err)
					return err
				}
				return nil
			})
		} else {
			if err = tx[0].Create(instance).Error; err != nil {
				logger.Errorln(err)
				return
			}
			userRole := &model.UserRole{
				BaseModel: common.BaseModel{ID: util.SnowflakeId()},
				UserID:    instance.ID,
				RoleID:    defaultRole.ID,
			}
			if err = tx[0].Create(userRole).Error; err != nil {
				logger.Errorln(err)
				return
			}
		}
	} else {
		if len(tx) == 0 {
			err = database.DB.Transaction(func(tx *gorm.DB) error {
				if err = tx.Create(instance).Error; err != nil {
					logger.Errorln(err)
					return err
				}
				return nil
			})
		} else {
			err = tx[0].Create(instance).Error
		}
	}
	if err != nil {
		logger.Errorln(err)
		return
	}
	// 完成后密码置空
	instance.Password = ""
	return
}

// UpdateDomain 更新用户信息
func (s *userService) UpdateDomain(instance *domain.UserDomain, tx ...*gorm.DB) (count int64, err error) {
	if lackedFields := instance.UpdateRequired(); lackedFields != "" {
		err = errors.New(lackedFields + " is required")
		return
	}
	if len(tx) > 0 {
		err = tx[0].Model(s.Model()).Where(instance.UpdateConditionModel()).Updates(instance.UpdateModel()).Error
	} else {
		err = database.DB.Transaction(func(t *gorm.DB) error {
			err = t.Model(s.Model()).Where(instance.UpdateConditionModel()).Updates(instance.UpdateModel()).Error
			if err != nil {
				logger.Errorln(err)
				return err
			}
			return nil
		})
	}
	if err != nil {
		logger.Errorln(err)
	}
	return
}

// UpdatePassword 更新用户密码
func (s *userService) UpdatePassword(instance *model.User) (success bool, err error) {
	if instance.ID == 0 {
		err = errors.New("id is required")
		return
	}
	if instance.Password == "" {
		err = errors.New("password is required")
		return
	}
	pwd, _ := bcrypt.GenerateFromPassword([]byte(instance.Password), bcrypt.DefaultCost)
	err = database.DB.Where(&model.User{BaseModel: common.BaseModel{ID: instance.ID}}).Updates(&model.User{
		Password: string(pwd),
	}).Error
	if err != nil {
		logger.Errorln(err)
		return
	}
	success = true
	return
}

// Instance 获取单个用户
func (s *userService) Instance(condition *model.User) (instance *model.User, err error) {
	if condition.ID == 0 {
		err = errors.New("id is required")
		return
	}
	instance = &model.User{}
	err = database.DB.Where(condition).First(instance).Error
	if err != nil {
		logger.Errorln(err)
		return
	}
	return
}

// Delete 删除用户
func (s *userService) Delete(condition *model.User, tx ...*gorm.DB) (count int64, err error) {
	if condition.ID == 0 {
		err = errors.New("id is required")
		return
	}

	var result *gorm.DB
	if len(tx) == 0 {
		result = database.DB.Delete(&model.User{}, condition)
	} else {
		result = tx[0].Delete(&model.User{}, condition)
	}
	if result.Error != nil {
		logger.Errorln(result.Error)
		return
	}
	count = result.RowsAffected
	return
}

// Roles 获取用户的角色列表
func (s *userService) Roles(userId uint64, types ...int) (roles []*model.Role, err error) {
	// 获取用户ID对应的所有角色信息
	sm := gorm.Statement{DB: database.DB}
	_ = sm.Parse(&model.Role{})
	ruleTableName := sm.Schema.Table

	tx := database.DB.Model(&model.UserRole{})
	if len(types) > 0 {
		tx = tx.Where("type in (?)", types)
	}
	err = tx.
		Select(ruleTableName + ".*").
		Joins("left join " + ruleTableName + " on " + ruleTableName + ".id = role_id").
		Where(&model.UserRole{UserID: userId}).Scan(&roles).Error

	//var list []*model.UserRole
	//err = database.DB.Where(&model.UserRole{UserID: userId}).Find(&list).Error
	//if err != nil {
	//	logger.Errorln(err)
	//	return
	//}
	//var roleIds []uint64
	//for _, v := range list {
	//	roleIds = append(roleIds, v.RoleID)
	//}

	return
}

func (s *userService) DispatchRoles(userID uint64, roleIDList []uint64) (success bool, err error) {
	if userID == 0 {
		err = errors.New("id is required")
		return
	}
	// 在事务中处理
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		// 删除原有的角色
		if err = tx.Where(&model.UserRole{UserID: userID}).Delete(&model.UserRole{}).Error; err != nil {
			logger.Errorln(err)
			return err
		}
		// 添加新的角色
		for _, v := range roleIDList {
			if err = tx.Create(&model.UserRole{
				BaseModel: common.BaseModel{ID: util.SnowflakeId()},
				UserID:    userID,
				RoleID:    v,
			}).Error; err != nil {
				logger.Errorln(err)
				return err
			}
		}
		return nil
	})

	if err != nil {
		return
	}

	success = true
	return
}
