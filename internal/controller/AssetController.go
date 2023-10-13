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

type assetController struct {
	common.BaseController[*model.Asset, *domain.AssetDomain]
}

func (c *assetController) InitManage() {
	r := server.Group("/api/v1/asset")

	r.Post("/add", middleware.NeedAuthorization(constant.ResourceAssetAdd), c.Add)
	r.Put("/update", middleware.NeedAuthorization(constant.ResourceAssetUpdate), c.Update)
	r.Post("/update", middleware.NeedAuthorization(constant.ResourceAssetUpdate), c.Update)
	r.Delete("/delete", middleware.NeedAuthorization(constant.ResourceAssetDelete), c.Delete)
	r.Post("/delete", middleware.NeedAuthorization(constant.ResourceAssetDelete), c.Delete)
	r.Get("/detail", middleware.NeedAuthorization(constant.ResourceAssetDetail), c.Detail)
	r.Get("/list", middleware.NeedAuthorization(constant.ResourceAssetList), c.List)
}

func (*assetController) GetService() common.Service[*model.Asset] {
	return service.AssetService
}

func (*assetController) NewModel() *model.Asset {
	return new(model.Asset)
}

func (*assetController) NewDomain() *domain.AssetDomain {
	return new(domain.AssetDomain)
}

func init() {
	c := new(assetController)
	c.BaseController = common.BaseController[*model.Asset, *domain.AssetDomain]{
		Controller: c,
	}

	Controllers = append(Controllers, c)
}
