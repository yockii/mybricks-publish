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

type pageVersionController struct {
	common.BaseController[*model.PageVersion, *domain.PageVersionDomain]
}

func (c *pageVersionController) InitManage() {
	r := server.Group("/api/v1/pageVersion")

	r.Post("/add", middleware.NeedAuthorization(constant.ResourcePageVersionAdd), c.Add)
	r.Put("/update", middleware.NeedAuthorization(constant.ResourcePageVersionUpdate), c.Update)
	r.Post("/update", middleware.NeedAuthorization(constant.ResourcePageVersionUpdate), c.Update)
	r.Delete("/delete", middleware.NeedAuthorization(constant.ResourcePageVersionDelete), c.Delete)
	r.Post("/delete", middleware.NeedAuthorization(constant.ResourcePageVersionDelete), c.Delete)
	r.Get("/detail", middleware.NeedAuthorization(constant.ResourcePageVersionDetail), c.Detail)
	r.Get("/list", middleware.NeedAuthorization(constant.ResourcePageVersionList), c.List)
}

func (*pageVersionController) GetService() common.Service[*model.PageVersion] {
	return service.PageVersionService
}

func (*pageVersionController) NewModel() *model.PageVersion {
	return new(model.PageVersion)
}

func (*pageVersionController) NewDomain() *domain.PageVersionDomain {
	return new(domain.PageVersionDomain)
}

func init() {
	c := new(pageVersionController)
	c.BaseController = common.BaseController[*model.PageVersion, *domain.PageVersionDomain]{
		Controller: c,
	}

	Controllers = append(Controllers, c)
}
