package model

import (
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"manatee-publish/pkg/common"
)

type Application struct {
	common.BaseModel
	Name        string         `json:"name" gorm:"size:100;comment:应用名称"`
	Description string         `json:"description" gorm:"size:1000;comment:应用描述"`
	CreateTime  int64          `json:"createTime" gorm:"autoCreateTime:milli"`
	UpdateTime  int64          `json:"updateTime" gorm:"autoUpdateTime:milli"`
	DeleteTime  gorm.DeletedAt `json:"deleteTime,omitempty" gorm:"index"`
}

func (*Application) TableComment() string {
	return `应用表`
}

func (m *Application) UnmarshalJSON(b []byte) error {
	j := gjson.ParseBytes(b)
	m.ID = j.Get("id").Uint()
	m.Name = j.Get("name").String()
	m.Description = j.Get("description").String()
	m.CreateTime = j.Get("createTime").Int()
	m.UpdateTime = j.Get("updateTime").Int()
	return nil
}

func (m *Application) AddRequired() string {
	if m.Name == "" {
		return "name"
	}
	return ""
}

func (m *Application) CheckDuplicatedModel() common.Model {
	return &Application{
		Name: m.Name,
	}
}
func (u *Application) UpdateModel() common.Model {
	return &Application{
		Name:        u.Name,
		Description: u.Description,
	}
}

func (u *Application) FuzzyQueryMap() map[string]string {
	return map[string]string{
		"name": "%" + u.Name + "%",
	}
}

// ApplicationPage 应用与页面关联
type ApplicationPage struct {
	common.BaseModel
	ApplicationID uint64 `json:"applicationId,omitempty,string" gorm:"index;comment:应用ID"`
	PageID        uint64 `json:"pageId,omitempty,string" gorm:"index;comment:页面ID"`
	CreateTime    int64  `json:"createTime" gorm:"autoCreateTime:milli"`
}

func (*ApplicationPage) TableComment() string {
	return `应用与页面关联表`
}

func (m *ApplicationPage) UnmarshalJSON(b []byte) error {
	j := gjson.ParseBytes(b)
	m.ID = j.Get("id").Uint()
	m.ApplicationID = j.Get("applicationId").Uint()
	m.PageID = j.Get("pageId").Uint()
	m.CreateTime = j.Get("createTime").Int()
	return nil
}

func (m *ApplicationPage) AddRequired() string {
	if m.ApplicationID == 0 || m.PageID == 0 {
		return "applicationId, pageId"
	}
	return ""
}

func (m *ApplicationPage) CheckDuplicatedModel() common.Model {
	return &ApplicationPage{
		ApplicationID: m.ApplicationID,
		PageID:        m.PageID,
	}
}

func init() {
	common.Models = append(common.Models, &Application{}, &ApplicationPage{})
}
