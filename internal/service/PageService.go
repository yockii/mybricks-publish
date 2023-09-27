package service

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/common"
)

var PageService = newPageService()

type pageService struct {
	common.BaseService[*model.Page]
}

func newPageService() *pageService {
	s := new(pageService)
	s.BaseService = common.BaseService[*model.Page]{
		Service: s,
	}
	return s
}

func (*pageService) Model() *model.Page {
	return new(model.Page)
}
