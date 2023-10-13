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

type assetVersionController struct {
	common.BaseController[*model.AssetVersion, *domain.AssetVersionDomain]
}

func (c *assetVersionController) InitManage() {
	r := server.Group("/api/v1/assetVersion")

	r.Post("/add", middleware.NeedAuthorization(constant.ResourceAssetVersionAdd), c.Add)
	r.Put("/update", middleware.NeedAuthorization(constant.ResourceAssetVersionUpdate), c.Update)
	r.Post("/update", middleware.NeedAuthorization(constant.ResourceAssetVersionUpdate), c.Update)
	r.Delete("/delete", middleware.NeedAuthorization(constant.ResourceAssetVersionDelete), c.Delete)
	r.Post("/delete", middleware.NeedAuthorization(constant.ResourceAssetVersionDelete), c.Delete)
	r.Get("/detail", middleware.NeedAuthorization(constant.ResourceAssetVersionDetail), c.Detail)
	r.Get("/list", middleware.NeedAuthorization(constant.ResourceAssetVersionList), c.List)
}

func (*assetVersionController) GetService() common.Service[*model.AssetVersion] {
	return service.AssetVersionService
}

func (*assetVersionController) NewModel() *model.AssetVersion {
	return new(model.AssetVersion)
}

func (*assetVersionController) NewDomain() *domain.AssetVersionDomain {
	return new(domain.AssetVersionDomain)
}

func init() {
	c := new(assetVersionController)
	c.BaseController = common.BaseController[*model.AssetVersion, *domain.AssetVersionDomain]{
		Controller: c,
	}

	Controllers = append(Controllers, c)
}
