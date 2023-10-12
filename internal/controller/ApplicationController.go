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

type appController struct {
	common.BaseController[*model.Application, *domain.ApplicationDomain]
}

func (c *appController) InitManage() {
	r := server.Group("/api/v1/app")

	r.Post("/add", middleware.NeedAuthorization(constant.ResourceApplicationAdd), c.Add)
	r.Put("/update", middleware.NeedAuthorization(constant.ResourceApplicationUpdate), c.Update)
	r.Post("/update", middleware.NeedAuthorization(constant.ResourceApplicationUpdate), c.Update)
	r.Delete("/delete", middleware.NeedAuthorization(constant.ResourceApplicationDelete), c.Delete)
	r.Post("/delete", middleware.NeedAuthorization(constant.ResourceApplicationDelete), c.Delete)
	r.Get("/detail", middleware.NeedAuthorization(constant.ResourceApplicationDetail), c.Detail)
	r.Get("/list", middleware.NeedAuthorization(constant.ResourceApplicationList), c.List)
}

func (*appController) GetService() common.Service[*model.Application] {
	return service.ApplicationService
}

func (*appController) NewModel() *model.Application {
	return new(model.Application)
}

func (*appController) NewDomain() *domain.ApplicationDomain {
	return new(domain.ApplicationDomain)
}

func init() {
	c := new(appController)
	c.BaseController = common.BaseController[*model.Application, *domain.ApplicationDomain]{
		Controller: c,
	}

	Controllers = append(Controllers, c)
}
