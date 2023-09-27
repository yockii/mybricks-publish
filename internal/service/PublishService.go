package service

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/common"
)

var PublishService = newPublishService()

type publishService struct {
	common.BaseService[*model.PublishInfo]
}

func newPublishService() *publishService {
	s := new(publishService)
	s.BaseService = common.BaseService[*model.PublishInfo]{
		Service: s,
	}
	return s
}

func (*publishService) Model() *model.PublishInfo {
	return new(model.PublishInfo)
}
