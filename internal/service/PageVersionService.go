package service

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/common"
)

var PageVersionService = newPageVersionService()

type pageVersionService struct {
	common.BaseService[*model.PageVersion]
}

func newPageVersionService() *pageVersionService {
	s := new(pageVersionService)
	s.BaseService = common.BaseService[*model.PageVersion]{
		Service: s,
	}
	return s
}

func (*pageVersionService) Model() *model.PageVersion {
	return new(model.PageVersion)
}
