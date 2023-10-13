package domain

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/server"
)

type PageVersionDomain struct {
	model.PageVersion
	CreateTimeCondition *server.TimeCondition `json:"createTimeCondition,omitempty"`
	OrderBy             string                `json:"orderBy,omitempty"`
}

func (r *PageVersionDomain) GetModel() *model.PageVersion {
	return &r.PageVersion
}

func (r *PageVersionDomain) GetOrderBy() string {
	return r.OrderBy
}

func (r *PageVersionDomain) GetTimeConditionList() map[string]*server.TimeCondition {
	return map[string]*server.TimeCondition{
		"create_time": r.CreateTimeCondition,
	}
}
