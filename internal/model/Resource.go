package model

import (
	"github.com/tidwall/gjson"
	"manatee-publish/pkg/common"
	"manatee-publish/pkg/util"
)

type Resource struct {
	common.BaseModel
	ResourceName string `json:"resourceName,omitempty" gorm:"size:50;comment:资源名称"`             // 资源名称
	ResourceCode string `json:"resourceCode,omitempty" gorm:"size:50;uniqueIndex;comment:资源代码"` // 资源认证代码
	Type         int    `json:"type,omitempty" gorm:"comment:类型 1-通用权限 2-专属权限"`                 // 类型
	CreateTime   int64  `json:"createTime" gorm:"autoCreateTime:milli"`
	UpdateTime   int64  `json:"updateTime" gorm:"autoUpdateTime:milli"`
}

func (_ *Resource) TableComment() string {
	return "资源表"
}
func (m *Resource) AddRequired() string {
	if m.ResourceName == "" || m.ResourceCode == "" || m.Type == 0 {
		return "resourceName,resourceCode,type"
	}
	return ""
}
func (m *Resource) UpdateModel() common.Model {
	return &Resource{
		ResourceName: m.ResourceName,
		ResourceCode: m.ResourceCode,
		Type:         m.Type,
	}
}
func (m *Resource) InitDefaultFields() {
	m.ID = util.SnowflakeId()
}
func (m *Resource) FuzzyQueryMap() map[string]string {
	result := make(map[string]string)
	if m.ResourceName != "" {
		result["resource_name"] = "%" + m.ResourceName + "%"
	}
	if m.ResourceCode != "" {
		result["resource_code"] = "%" + m.ResourceCode + "%"
	}
	return result
}
func (m *Resource) ExactMatchModel() common.Model {
	b := new(Resource)
	b.Type = m.Type
	return b
}

type RoleResource struct {
	common.BaseModel
	RoleID       uint64 `json:"roleId,omitempty,string"`
	ResourceCode string `json:"resourceCode,omitempty" gorm:"size:50;comment:资源代码"` // 资源认证代码
	CreateTime   int64  `json:"createTime" gorm:"autoCreateTime:milli"`
}

func (_ *RoleResource) TableComment() string {
	return "角色资源表"
}
func (rr *RoleResource) UnmarshalJSON(b []byte) error {
	j := gjson.ParseBytes(b)
	rr.ID = j.Get("id").Uint()
	rr.RoleID = j.Get("roleId").Uint()
	rr.ResourceCode = j.Get("resourceCode").String()
	return nil
}

func init() {
	common.Models = append(common.Models, &Resource{}, &RoleResource{})
}
