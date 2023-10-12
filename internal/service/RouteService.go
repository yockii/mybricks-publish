package service

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/common"
)

var RouteService = newRouteService()

type routeService struct {
	common.BaseService[*model.Route]
}

func newRouteService() *routeService {
	s := new(routeService)
	s.BaseService = common.BaseService[*model.Route]{
		Service: s,
	}
	return s
}

func (*routeService) Model() *model.Route {
	return new(model.Route)
}
