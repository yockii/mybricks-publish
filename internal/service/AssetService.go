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
func (s *assetService) AddAsset(instance *model.Asset, reader io.Reader) (assetVersionID uint64, err error) {
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
		if err = tx.Create(instance).Error; err != nil {
			logger.Errorln(err)
			return err
		}

		// 添加文件版本信息
		if err = tx.Create(&model.AssetVersion{
			BaseModel: common.BaseModel{
				ID: assetVersionID,
			},
			FileID:      instance.ID,
			OssConfigID: s.osManager.GetOssConfigID(),
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

func (s *assetService) Download(path string) (reader io.ReadCloser, err error) {
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
	return s.osManager.GetObject(instance.ObjName)
}
