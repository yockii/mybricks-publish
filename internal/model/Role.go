package model

import (
	"github.com/tidwall/gjson"
	"manatee-publish/pkg/common"
	"manatee-publish/pkg/util"
)

const (
	RoleTypeSuperAdmin = -1
	RoleTypeNormal     = 1
)

const (
	RoleDataPermissionAll        = 1
	RoleDataPermissionDept       = 2
	RoleDataPermissionDeptAndSub = 3
	RoleDataPermissionSelf       = 4
)

const (
	RoleStatusNormal  = 1
	RoleStatusDisable = 2
)

type Role struct {
	common.BaseModel
	Name           string `json:"name,omitempty" gorm:"size:50;comment:角色名称"`
	Desc           string `json:"desc,omitempty" gorm:"size:200;comment:角色描述"`
	Type           int    `json:"type,omitempty" gorm:"comment:角色类型 1-普通用户角色 2-家政服务人员角色 3-系统运维人员角色 -1-超级管理员角色"`
	DataPermission int    `json:"dataPermission,omitempty" gorm:"comment:数据权限 1-全部数据权限 2-本部门及以下数据权限 3-仅本人数据权限"`
	Style          string `json:"style,omitempty" gorm:"size:500;comment:角色样式"`
	DefaultRole    int    `json:"defaultRole" gorm:"comment:默认角色 1-是 其他否"`
	Status         int    `json:"status,omitempty" gorm:"comment:状态 1-启用 2-禁用"`
	CreateTime     int64  `json:"createTime" gorm:"autoCreateTime:milli"`
	UpdateTime     int64  `json:"updateTime" gorm:"autoUpdateTime:milli"`
}

func (_ *Role) TableComment() string {
	return "角色表"
}

func (m *Role) AddRequired() string {
	if m.Name == "" || m.Type == 0 {
		return "name,type"
	}
	return ""
}
func (m *Role) UpdateModel() common.Model {
	return &Role{
		Name:           m.Name,
		Desc:           m.Desc,
		Type:           m.Type,
		DataPermission: m.DataPermission,
		Style:          m.Style,
		Status:         m.Status,
	}
}
func (m *Role) InitDefaultFields() {
	m.ID = util.SnowflakeId()
	m.Status = RoleStatusNormal
	if m.DataPermission == 0 {
		m.DataPermission = 3
	}
}
func (m *Role) FuzzyQueryMap() map[string]string {
	result := make(map[string]string)
	if m.Name != "" {
		result["name"] = "%" + m.Name + "%"
	}
	return result
}
func (m *Role) ExactMatchModel() common.Model {
	b := new(Role)
	b.Type = m.Type
	b.Status = m.Status
	return b
}

type UserRole struct {
	common.BaseModel
	UserID     uint64 `json:"userId,omitempty,string"`
	RoleID     uint64 `json:"roleId,omitempty,string"`
	CreateTime int64  `json:"createTime" gorm:"autoCreateTime:milli"`
}

func (_ *UserRole) TableComment() string {
	return "用户角色表"
}
func (ur *UserRole) UnmarshalJSON(b []byte) error {
	j := gjson.ParseBytes(b)
	ur.ID = j.Get("id").Uint()
	ur.UserID = j.Get("userId").Uint()
	ur.RoleID = j.Get("roleId").Uint()
	return nil
}

func init() {
	common.Models = append(common.Models, &Role{}, &UserRole{})
}
