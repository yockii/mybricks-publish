package model

import (
	"github.com/tidwall/gjson"
	"manatee-publish/pkg/common"
)

// Route 本应用自身页面管理
type Route struct {
	common.BaseModel
	Code        string `gorm:"size:255;not null;comment:路由编码" json:"code,omitempty"`
	DisplayName string `gorm:"size:255;not null;comment:路由名称" json:"displayName,omitempty"`
	ActiveRule  string `gorm:"size:255;not null;comment:路由路径" json:"activeRule,omitempty"`
	Entry       string `gorm:"size:500;not null;comment:路由加载地址" json:"entry,omitempty"`
	CreateTime  int64  `json:"createTime" gorm:"autoCreateTime:milli" json:"createTime,omitempty"`
}

func (*Route) TableComment() string {
	return `路由表`
}

func (m *Route) UnmarshalJSON(b []byte) error {
	j := gjson.ParseBytes(b)
	m.ID = j.Get("id").Uint()
	m.Code = j.Get("code").String()
	m.DisplayName = j.Get("displayName").String()
	m.ActiveRule = j.Get("activeRule").String()
	m.Entry = j.Get("entry").String()
	m.CreateTime = j.Get("createTime").Int()
	return nil
}

func (m *Route) AddRequired() string {
	if m.Code == "" || m.DisplayName == "" || m.ActiveRule == "" || m.Entry == "" {
		return "code,name,activeRule,entry"
	}
	return ""
}

func (m *Route) CheckDuplicatedModel() common.Model {
	return &Route{
		Code: m.Code,
	}
}

func (u *Route) UpdateModel() common.Model {
	return &Route{
		Code:        u.Code,
		DisplayName: u.DisplayName,
		ActiveRule:  u.ActiveRule,
		Entry:       u.Entry,
	}
}

func (u *Route) FuzzyQueryMap() map[string]string {
	result := make(map[string]string)
	if u.Code != "" {
		result["code"] = "%" + u.Code + "%"
	}
	if u.DisplayName != "" {
		result["display_name"] = "%" + u.DisplayName + "%"
	}
	if u.ActiveRule != "" {
		result["active_rule"] = "%" + u.ActiveRule + "%"
	}
	if u.Entry != "" {
		result["entry"] = "%" + u.Entry + "%"
	}
	return result
}

func init() {
	common.Models = append(common.Models, &Route{})
}
