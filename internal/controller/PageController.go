package controller

import (
	"manatee-publish/internal/constant"
	"manatee-publish/internal/domain"
	"manatee-publish/internal/middleware"
	"manatee-publish/internal/model"
	"manatee-publish/internal/service"
	"manatee-publish/pkg/common"
	"manatee-publish/pkg/server"
)

type pageController struct {
	common.BaseController[*model.Page, *domain.PageDomain]
}

func (c *pageController) InitManage() {
	r := server.Group("/api/v1/page")

	r.Post("/add", middleware.NeedAuthorization(constant.ResourcePageAdd), c.Add)
	r.Put("/update", middleware.NeedAuthorization(constant.ResourcePageUpdate), c.Update)
	r.Post("/update", middleware.NeedAuthorization(constant.ResourcePageUpdate), c.Update)
	r.Delete("/delete", middleware.NeedAuthorization(constant.ResourcePageDelete), c.Delete)
	r.Post("/delete", middleware.NeedAuthorization(constant.ResourcePageDelete), c.Delete)
	r.Get("/detail", middleware.NeedAuthorization(constant.ResourcePageDetail), c.Detail)
	r.Get("/list", middleware.NeedAuthorization(constant.ResourcePageList), c.List)
}

func (*pageController) GetService() common.Service[*model.Page] {
	return service.PageService
}

func (*pageController) NewModel() *model.Page {
	return new(model.Page)
}

func (*pageController) NewDomain() *domain.PageDomain {
	return new(domain.PageDomain)
}

func init() {
	c := new(pageController)
	c.BaseController = common.BaseController[*model.Page, *domain.PageDomain]{
		Controller: c,
	}

	Controllers = append(Controllers, c)
}
