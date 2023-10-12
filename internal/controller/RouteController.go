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

type routeController struct {
	common.BaseController[*model.Route, *domain.RouteDomain]
}

func (c *routeController) InitManage() {
	r := server.Group("/api/v1/route")

	r.Post("/add", middleware.NeedAuthorization(constant.ResourceRouteAdd), c.Add)
	r.Put("/update", middleware.NeedAuthorization(constant.ResourceRouteUpdate), c.Update)
	r.Post("/update", middleware.NeedAuthorization(constant.ResourceRouteUpdate), c.Update)
	r.Delete("/delete", middleware.NeedAuthorization(constant.ResourceRouteDelete), c.Delete)
	r.Post("/delete", middleware.NeedAuthorization(constant.ResourceRouteDelete), c.Delete)
	r.Get("/detail", middleware.NeedAuthorization(constant.ResourceRouteDetail), c.Detail)
	r.Get("/list", middleware.NeedAuthorization(constant.ResourceRouteList), c.List)
}

func (*routeController) GetService() common.Service[*model.Route] {
	return service.RouteService
}

func (*routeController) NewModel() *model.Route {
	return new(model.Route)
}

func (*routeController) NewDomain() *domain.RouteDomain {
	return new(domain.RouteDomain)
}

func init() {
	c := new(routeController)
	c.BaseController = common.BaseController[*model.Route, *domain.RouteDomain]{
		Controller: c,
	}

	Controllers = append(Controllers, c)
}
