package domain

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/server"
)

type OssConfigDomain struct {
	model.OssConfig
	CreateTimeCondition *server.TimeCondition `json:"createTimeCondition,omitempty"`
	OrderBy             string                `json:"orderBy,omitempty"`
}

func (r *OssConfigDomain) GetModel() *model.OssConfig {
	return &r.OssConfig
}

func (r *OssConfigDomain) GetOrderBy() string {
	return r.OrderBy
}

func (r *OssConfigDomain) GetTimeConditionList() map[string]*server.TimeCondition {
	return map[string]*server.TimeCondition{
		"create_time": r.CreateTimeCondition,
	}
}
