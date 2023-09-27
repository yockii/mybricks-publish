package model

import (
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"manatee-publish/pkg/common"
)

type Asset struct {
	common.BaseModel
	OssConfigID uint64         `json:"ossConfigId,omitempty,string" gorm:"comment:OSS配置ID"`
	Path        string         `json:"path" gorm:"size:1000;comment:路径"`
	ObjName     string         `json:"objName,omitempty" gorm:"size:200;comment:存储的对象名称"`
	CreateTime  int64          `json:"createTime" gorm:"autoCreateTime:milli"`
	UpdateTime  int64          `json:"updateTime" gorm:"autoUpdateTime:milli"`
	DeleteTime  gorm.DeletedAt `json:"deleteTime,omitempty" gorm:"index"`
}

func (*Asset) TableComment() string {
	return `资产文件表`
}

func (af *Asset) UnmarshalJSON(b []byte) error {
	j := gjson.ParseBytes(b)
	af.ID = j.Get("id").Uint()
	af.OssConfigID = j.Get("ossConfigId").Uint()
	af.Path = j.Get("path").String()
	af.ObjName = j.Get("objName").String()
	af.CreateTime = j.Get("createTime").Int()
	af.UpdateTime = j.Get("updateTime").Int()
	return nil
}

func init() {
	common.Models = append(common.Models, &Asset{})
}
