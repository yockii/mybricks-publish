package domain

import (
	"manatee-publish/internal/model"
	"manatee-publish/pkg/server"
)

type PublishInfoDomain struct {
	model.Page
	Content PublishContent `json:"content"`

	CreateTimeCondition *server.TimeCondition `json:"createTimeCondition"`
	UpdateTimeCondition *server.TimeCondition `json:"updateTimeCondition"`
	OrderBy             string                `json:"orderBy"`
}

type PublishContent struct {
	Json        string              `json:"json"`
	Html        string              `json:"html"`
	Js          []ContentJs         `json:"js"`
	Permissions []ContentPermission `json:"permissions"`
	GlobalDeps  []GlobalDep         `json:"globalDeps"`
	Images      []ContentImage      `json:"images"`
}

type ContentImage struct {
	Path    string `json:"path,omitempty"`
	Content string `json:"content,omitempty"`
}

type GlobalDep struct {
	Path    string `json:"path,omitempty"`
	Content string `json:"content,omitempty"`
}

type ContentJs struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
}

type ContentPermission struct {
	Code   string `json:"code,omitempty"`
	Title  string `json:"title,omitempty"`
	Remark string `json:"remark,omitempty"`
}

func (r *PublishInfoDomain) GetModel() *model.Page {
	return &r.Page
}

func (r *PublishInfoDomain) GetOrderBy() string {
	return r.OrderBy
}

func (r *PublishInfoDomain) GetTimeConditionList() map[string]*server.TimeCondition {
	return map[string]*server.TimeCondition{
		"create_time": r.CreateTimeCondition,
		"update_time": r.UpdateTimeCondition,
	}
}

type ManateePublishResponse struct {
	Code    int         `json:"code,omitempty"` // 成功1，失败0
	Data    PublishData `json:"data"`
	Message string      `json:"message,omitempty"` // 自定义内容
}

type PublishData struct {
	Url string `json:"url,omitempty"`
}
