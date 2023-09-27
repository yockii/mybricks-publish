package service

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/common"
)

var ResourceService = newResourceService()

type resourceService struct {
	common.BaseService[*model.Resource]
}

func newResourceService() *resourceService {
	s := new(resourceService)
	s.BaseService = common.BaseService[*model.Resource]{
		Service: s,
	}
	return s
}

func (*resourceService) Model() *model.Resource {
	return new(model.Resource)
}
