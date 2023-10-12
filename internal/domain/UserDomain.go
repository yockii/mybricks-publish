package domain

import (
	"github.com/tidwall/gjson"
	"manatee-publish/internal/model"
	"manatee-publish/pkg/server"
)

type UserDomain struct {
	model.User
	OrderBy             string                `json:"orderBy,omitempty"`
	CreateTimeCondition *server.TimeCondition `json:"createTimeCondition,omitempty"`
}

func (d *UserDomain) GetOrderBy() string {
	return d.OrderBy
}

func (d *UserDomain) GetModel() *model.User {
	return &d.User
}

func (d *UserDomain) GetTimeConditionList() map[string]*server.TimeCondition {
	return map[string]*server.TimeCondition{
		"create_time": d.CreateTimeCondition,
	}
}

type UserDispatchRolesRequest struct {
	UserID     uint64   `json:"userId,string"`
	RoleIDList []uint64 `json:"roleIdList"`
}

func (r *UserDispatchRolesRequest) UnmarshalJSON(b []byte) error {
	j := gjson.ParseBytes(b)
	r.UserID = j.Get("userId").Uint()
	for _, v := range j.Get("roleIdList").Array() {
		r.RoleIDList = append(r.RoleIDList, v.Uint())
	}
	return nil
}
