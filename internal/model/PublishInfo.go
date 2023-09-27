package model

import "manatee-publish/pkg/common"

type PublishInfo struct {
	common.BaseModel
	ProductID      int64  `json:"productID,omitempty" gorm:"comment:页面ID"`
	ProductName    string `json:"productName,omitempty" gorm:"size:255;comment:页面名称"`
	Env            string `json:"env,omitempty" gorm:"size:50;comment:发布环境标识"`
	PublisherEmail string `json:"publisherEmail,omitempty" gorm:"size:50;comment:发布人邮箱"`
	PublisherName  string `json:"publisherName,omitempty" gorm:"size:50;comment:发布人名称"`
	Version        string `json:"version,omitempty" gorm:"size:50;comment:发布版本号"`
	Type           string `json:"type,omitempty" gorm:"size:20;comment:页面类型"`
	GroupID        int64  `json:"groupID,omitempty" gorm:"comment:协作组ID"`
	GroupName      string `json:"groupName,omitempty" gorm:"size:200;comment:协作组名称"`
	CommitInfo     string `json:"commitInfo,omitempty" gorm:"size:500;comment:发布提交信息"`
	SchemaJson     string `json:"schemaJson,omitempty" gorm:"comment:页面搭建产物schema"`
	CreateTime     int64  `json:"createTime" gorm:"autoCreateTime:milli"`
	UpdateTime     int64  `json:"updateTime" gorm:"autoUpdateTime:milli"`
}

func (_ *PublishInfo) TableComment() string {
	return "发布信息表"
}

func (m *PublishInfo) FuzzyQueryMap() map[string]string {
	result := make(map[string]string)
	if m.ProductName != "" {
		result["product_name"] = "%" + m.ProductName + "%"
	}
	if m.PublisherName != "" {
		result["publisher_name"] = "%" + m.PublisherName + "%"
	}
	if m.GroupName != "" {
		result["group_name"] = "%" + m.GroupName + "%"
	}
	return result
}

func (m *PublishInfo) ExactMatchModel() common.Model {
	b := new(PublishInfo)
	b.Type = m.Type
	b.GroupID = m.GroupID
	return b
}

func init() {
	common.Models = append(common.Models, &PublishInfo{})
}
