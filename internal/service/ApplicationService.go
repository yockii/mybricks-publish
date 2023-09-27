package service

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/common"
)

var ApplicationService = newApplicationService()

type applicationService struct {
	common.BaseService[*model.Application]
}

func newApplicationService() *applicationService {
	s := new(applicationService)
	s.BaseService = common.BaseService[*model.Application]{
		Service: s,
	}
	return s
}

func (*applicationService) Model() *model.Application {
	return new(model.Application)
}

func (*applicationService) CheckOrigin(pageId uint64, origin string) bool {
	return true
}
