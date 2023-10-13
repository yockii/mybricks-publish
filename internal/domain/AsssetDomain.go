package domain

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/server"
)

type AssetDomain struct {
	model.Asset
	CreateTimeCondition *server.TimeCondition `json:"createTimeCondition,omitempty"`
	OrderBy             string                `json:"orderBy,omitempty"`
}

func (r *AssetDomain) GetModel() *model.Asset {
	return &r.Asset
}

func (r *AssetDomain) GetOrderBy() string {
	return r.OrderBy
}

func (r *AssetDomain) GetTimeConditionList() map[string]*server.TimeCondition {
	return map[string]*server.TimeCondition{
		"create_time": r.CreateTimeCondition,
	}
}
