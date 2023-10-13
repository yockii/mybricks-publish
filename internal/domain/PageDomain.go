package domain

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/server"
)

type PageDomain struct {
	model.Page
	CreateTimeCondition *server.TimeCondition `json:"createTimeCondition,omitempty"`
	OrderBy             string                `json:"orderBy,omitempty"`
}

func (r *PageDomain) GetModel() *model.Page {
	return &r.Page
}

func (r *PageDomain) GetOrderBy() string {
	return r.OrderBy
}

func (r *PageDomain) GetTimeConditionList() map[string]*server.TimeCondition {
	return map[string]*server.TimeCondition{
		"create_time": r.CreateTimeCondition,
	}
}
