package domain

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/server"
)

type ApplicationDomain struct {
	model.Application
	CreateTimeCondition *server.TimeCondition `json:"createTimeCondition,omitempty"`
	OrderBy             string                `json:"orderBy,omitempty"`
}

func (r *ApplicationDomain) GetModel() *model.Application {
	return &r.Application
}

func (r *ApplicationDomain) GetOrderBy() string {
	return r.OrderBy
}

func (r *ApplicationDomain) GetTimeConditionList() map[string]*server.TimeCondition {
	return map[string]*server.TimeCondition{
		"create_time": r.CreateTimeCondition,
	}
}
