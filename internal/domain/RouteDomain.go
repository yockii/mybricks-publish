package domain

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/server"
)

type RouteDomain struct {
	model.Route
	CreateTimeCondition *server.TimeCondition `json:"createTimeCondition,omitempty"`
	OrderBy             string                `json:"orderBy,omitempty"`
}

func (r *RouteDomain) GetModel() *model.Route {
	return &r.Route
}

func (r *RouteDomain) GetOrderBy() string {
	return r.OrderBy
}

func (r *RouteDomain) GetTimeConditionList() map[string]*server.TimeCondition {
	return map[string]*server.TimeCondition{
		"create_time": r.CreateTimeCondition,
	}
}
