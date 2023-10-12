package service

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/common"
)

var OssConfigService = newOssConfigService()

type ossConfigService struct {
	common.BaseService[*model.OssConfig]
}

func newOssConfigService() *ossConfigService {
	s := new(ossConfigService)
	s.BaseService = common.BaseService[*model.OssConfig]{
		Service: s,
	}
	return s
}

func (*ossConfigService) Model() *model.OssConfig {
	return new(model.OssConfig)
}
