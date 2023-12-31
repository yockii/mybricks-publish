package service

import (
	"errors"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"manatee-publish/internal/model"
	provider "manatee-publish/internal/oss"
	"manatee-publish/pkg/common"
	"manatee-publish/pkg/database"
	"manatee-publish/pkg/util"
	"strings"
	"sync"
	"time"
)

var AssetService = &assetService{
	locker: new(sync.Mutex),
}

type assetService struct {
	common.BaseService[*model.Asset]
	osManager provider.OsManager
	locker    sync.Locker
}

func (s *assetService) initOsManager() (err error) {
	s.locker.Lock()
	defer s.locker.Unlock()
	if s.osManager == nil {
		// 从数据库中取出可用的云存储配置
		c := &model.OssConfig{}
		if err = database.DB.Where(&model.OssConfig{Status: 1}).First(c).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				err = errors.New("no available oss config")
			}
			return
		}
		// 初始化云存储
		osm := provider.GetProvider(c)
		if err = osm.Auth(); err != nil {
			return
		}
		s.osManager = osm
	}
	return nil
}

// AddAsset 添加资源
func (s *assetService) AddAsset(instance *model.Asset, version string, reader io.Reader) (assetVersionID uint64, err error) {
	if instance.Path == "" {
		err = errors.New("Path is required ")
		return
	}
	if s.osManager == nil {
		// 初始化新的osManager
		if err = s.initOsManager(); err != nil {
			return
		}
	}

	// 检查是否已有path重复，若有则取出
	var asset *model.Asset
	if err = database.DB.Where(&model.Asset{Path: instance.Path}).First(&asset).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			asset = nil
		} else {
			logger.Errorln(err)
			return
		}
	}

	suffix := instance.Path[strings.LastIndex(instance.Path, ".")+1:]

	instance.ID = util.SnowflakeId()
	// 上传文件
	now := time.Now().Format("20060102")
	objName := fmt.Sprintf("%s/%d.%s", now, instance.ID, suffix)
	if err = s.osManager.PutObject(objName, reader); err != nil {
		return
	}

	instance.OssConfigID = s.osManager.GetOssConfigID()
	instance.ObjName = objName

	assetVersionID = util.SnowflakeId()
	if err = database.DB.Transaction(func(tx *gorm.DB) error {
		assetId := instance.ID
		if asset == nil {
			if err = tx.Create(instance).Error; err != nil {
				logger.Errorln(err)
				return err
			}
		} else {
			assetId = asset.ID
			// 更新asset记录中的objName
			if err = tx.Model(&model.Asset{}).Where(&model.Asset{BaseModel: common.BaseModel{ID: assetId}}).Update("obj_name", objName).Error; err != nil {
				logger.Errorln(err)
				return err
			}
		}

		// 添加文件版本信息
		if err = tx.Create(&model.AssetVersion{
			BaseModel: common.BaseModel{
				ID: assetVersionID,
			},
			FileID:      assetId,
			OssConfigID: s.osManager.GetOssConfigID(),
			Version:     version,
			ObjName:     objName,
		}).Error; err != nil {
			logger.Errorln(err)
			return err
		}

		return nil
	}); err != nil {
		return
	}
	return
}

func (s *assetService) Count(path string) int64 {
	var count int64
	database.DB.Model(&model.Asset{}).Where(&model.Asset{Path: path}).Count(&count)
	return count
}

func (s *assetService) Download(path, version string) (reader io.ReadCloser, err error) {
	// 根据path获取文件信息
	instance := &model.Asset{}
	if err = database.DB.Where(&model.Asset{Path: path}).First(instance).Error; err != nil {
		logger.Errorln(err)
		return
	}
	if s.osManager == nil {
		// 初始化新的osManager
		if err = s.initOsManager(); err != nil {
			return
		}
	}
	if version == "" {
		return s.osManager.GetObject(instance.ObjName)
	} else {
		// 如果有版本，则要从assetVersion中获取
		assetVersion := &model.AssetVersion{}
		assetVersion.FileID = instance.ID
		assetVersion.Version = version
		if err = database.DB.Where(assetVersion).First(assetVersion).Error; err != nil {
			logger.Errorln(err)
			return
		}
		return s.osManager.GetObject(assetVersion.ObjName)
	}
}
