package domain

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/server"
)

type AssetVersionDomain struct {
	model.AssetVersion
	CreateTimeCondition *server.TimeCondition `json:"createTimeCondition,omitempty"`
	OrderBy             string                `json:"orderBy,omitempty"`
}

func (r *AssetVersionDomain) GetModel() *model.AssetVersion {
	return &r.AssetVersion
}

func (r *AssetVersionDomain) GetOrderBy() string {
	return r.OrderBy
}

func (r *AssetVersionDomain) GetTimeConditionList() map[string]*server.TimeCondition {
	return map[string]*server.TimeCondition{
		"create_time": r.CreateTimeCondition,
	}
}
