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

type ossConfigController struct {
	common.BaseController[*model.OssConfig, *domain.OssConfigDomain]
}

func (c *ossConfigController) InitManage() {
	r := server.Group("/api/v1/ossConfig")

	r.Post("/add", middleware.NeedAuthorization(constant.ResourceOssConfigAdd), c.Add)
	r.Put("/update", middleware.NeedAuthorization(constant.ResourceOssConfigUpdate), c.Update)
	r.Post("/update", middleware.NeedAuthorization(constant.ResourceOssConfigUpdate), c.Update)
	r.Delete("/delete", middleware.NeedAuthorization(constant.ResourceOssConfigDelete), c.Delete)
	r.Post("/delete", middleware.NeedAuthorization(constant.ResourceOssConfigDelete), c.Delete)
	r.Get("/detail", middleware.NeedAuthorization(constant.ResourceOssConfigDetail), c.Detail)
	r.Get("/list", middleware.NeedAuthorization(constant.ResourceOssConfigList), c.List)
}

func (*ossConfigController) GetService() common.Service[*model.OssConfig] {
	return service.OssConfigService
}

func (*ossConfigController) NewModel() *model.OssConfig {
	return new(model.OssConfig)
}

func (*ossConfigController) NewDomain() *domain.OssConfigDomain {
	return new(domain.OssConfigDomain)
}

func init() {
	c := new(ossConfigController)
	c.BaseController = common.BaseController[*model.OssConfig, *domain.OssConfigDomain]{
		Controller: c,
	}

	Controllers = append(Controllers, c)
}
