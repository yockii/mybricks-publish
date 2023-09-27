package domain

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/server"
)

type RoleDomain struct {
	model.Role
	CreateTimeCondition *server.TimeCondition `json:"createTimeCondition"`
	UpdateTimeCondition *server.TimeCondition `json:"updateTimeCondition"`
	OrderBy             string                `json:"orderBy"`
}

func (r *RoleDomain) GetModel() *model.Role {
	return &r.Role
}

func (r *RoleDomain) GetOrderBy() string {
	return r.OrderBy
}

func (r *RoleDomain) GetTimeConditionList() map[string]*server.TimeCondition {
	return map[string]*server.TimeCondition{
		"create_time": r.CreateTimeCondition,
		"update_time": r.UpdateTimeCondition,
	}
}

type RoleDispatchResourcesRequest struct {
	RoleID           uint64   `json:"roleId,string"`
	ResourceCodeList []string `json:"resourceCodeList"`
}
