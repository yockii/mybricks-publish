package model

import (
	"github.com/tidwall/gjson"
	"manatee-publish/pkg/common"
)

// Page 页面信息
type Page struct {
	common.BaseModel
	ProductID     int64  `json:"productId,omitempty" gorm:"comment:页面ID,对应传入的productId"`
	Name          string `json:"name,omitempty" gorm:"size:255;comment:页面名称"`
	GroupName     string `json:"groupName,omitempty" gorm:"size:255;comment:协作组组名称"`
	Type          string `json:"type,omitempty" gorm:"size:20;comment:页面类型"`
	PageVersionID uint64 `json:"pageVersionId,omitempty" gorm:"comment:页面版本ID"`
	CreateTime    int64  `json:"createTime" gorm:"autoCreateTime:milli"`
	UpdateTime    int64  `json:"updateTime" gorm:"autoUpdateTime:milli"`
}

func (_ *Page) TableComment() string {
	return "页面信息表"
}

func (m *Page) UnmarshalJSON(b []byte) error {
	j := gjson.ParseBytes(b)
	m.ID = j.Get("id").Uint()
	m.ProductID = j.Get("productId").Int()
	m.Name = j.Get("name").String()
	m.GroupName = j.Get("groupName").String()
	m.Type = j.Get("type").String()
	m.PageVersionID = j.Get("pageVersionId").Uint()
	m.CreateTime = j.Get("createTime").Int()
	m.UpdateTime = j.Get("updateTime").Int()
	return nil
}

func (m *Page) AddRequired() string {
	if m.ProductID == 0 || m.Name == "" {
		return "productId, name"
	}
	return ""
}

func (m *Page) CheckDuplicatedModel() common.Model {
	return &Page{
		ProductID: m.ProductID,
	}
}

func (m *Page) InstanceRequired() string {
	if m.ID == 0 && m.ProductID == 0 {
		return "id or product"
	}
	return ""
}

func (m *Page) UpdateModel() common.Model {
	return &Page{
		ProductID:     m.ProductID,
		Name:          m.Name,
		GroupName:     m.GroupName,
		Type:          m.Type,
		PageVersionID: m.PageVersionID,
	}
}

func (m *Page) FuzzyQueryMap() map[string]string {
	result := make(map[string]string)
	if m.Name != "" {
		result["name"] = "%" + m.Name + "%"
	}
	if m.GroupName != "" {
		result["group_name"] = "%" + m.GroupName + "%"
	}
	return result
}

func (m *Page) ExactMatchModel() common.Model {
	b := new(Page)
	b.ProductID = m.ProductID
	b.Type = m.Type
	b.PageVersionID = m.PageVersionID
	return b
}

// PageVersion 页面版本信息
type PageVersion struct {
	common.BaseModel
	PageID         uint64 `json:"pageId,omitempty" gorm:"comment:页面ID"`
	Version        string `json:"version,omitempty" gorm:"size:50;comment:版本号"`
	Env            string `json:"env,omitempty" gorm:"size:50;comment:发布环境标识"`
	PublisherEmail string `json:"publisherEmail,omitempty" gorm:"size:50;comment:发布人邮箱"`
	PublisherName  string `json:"publisherName,omitempty" gorm:"size:50;comment:发布人名称"`
	CommitInfo     string `json:"commitInfo,omitempty" gorm:"size:500;comment:发布提交信息"`
	SchemaJson     string `json:"schemaJson,omitempty" gorm:"comment:页面搭建产物schema"`
	AssetVersionID uint64 `json:"assetVersionId,omitempty" gorm:"comment:资产版本ID"`
	CreateTime     int64  `json:"createTime" gorm:"autoCreateTime:milli"`
	UpdateTime     int64  `json:"updateTime" gorm:"autoUpdateTime:milli"`
}

func (_ *PageVersion) TableComment() string {
	return "页面版本信息表"
}

func (m *PageVersion) UnmarshalJSON(b []byte) error {
	j := gjson.ParseBytes(b)
	m.ID = j.Get("id").Uint()
	m.PageID = j.Get("pageId").Uint()
	m.Version = j.Get("version").String()
	m.Env = j.Get("env").String()
	m.PublisherEmail = j.Get("publisherEmail").String()
	m.PublisherName = j.Get("publisherName").String()
	m.CommitInfo = j.Get("commitInfo").String()
	m.SchemaJson = j.Get("schemaJson").String()
	m.AssetVersionID = j.Get("assetVersionId").Uint()
	m.CreateTime = j.Get("createTime").Int()
	m.UpdateTime = j.Get("updateTime").Int()
	return nil
}

func (m *PageVersion) AddRequired() string {
	if m.PageID == 0 || m.Version == "" || m.Env == "" {
		return "pageId, version, env"
	}
	return ""
}

func (m *PageVersion) CheckDuplicatedModel() common.Model {
	return &PageVersion{
		PageID:  m.PageID,
		Version: m.Version,
		Env:     m.Env,
	}
}

func (m *PageVersion) UpdateModel() common.Model {
	return &PageVersion{
		PageID:         m.PageID,
		Version:        m.Version,
		Env:            m.Env,
		PublisherEmail: m.PublisherEmail,
		PublisherName:  m.PublisherName,
		CommitInfo:     m.CommitInfo,
		SchemaJson:     m.SchemaJson,
		AssetVersionID: m.AssetVersionID,
	}
}

func (m *PageVersion) FuzzyQueryMap() map[string]string {
	result := make(map[string]string)
	if m.Version != "" {
		result["version"] = "%" + m.Version + "%"
	}
	if m.Env != "" {
		result["env"] = "%" + m.Env + "%"
	}
	if m.PublisherEmail != "" {
		result["publisher_email"] = "%" + m.PublisherEmail + "%"
	}
	if m.PublisherName != "" {
		result["publisher_name"] = "%" + m.PublisherName + "%"
	}
	return result
}

func (m *PageVersion) ExactMatchModel() common.Model {
	b := new(PageVersion)
	b.PageID = m.PageID
	b.Version = m.Version
	b.Env = m.Env
	b.AssetVersionID = m.AssetVersionID
	return b
}

func init() {
	common.Models = append(common.Models, &Page{}, &PageVersion{})
}
