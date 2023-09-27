package model

import (
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"manatee-publish/pkg/common"
)

type AssetVersion struct {
	common.BaseModel
	FileID      uint64         `json:"fileId,omitempty,string" gorm:"index;comment:文件ID"`
	OssConfigID uint64         `json:"ossConfigId,omitempty,string" gorm:"comment:OSS配置ID"`
	ObjName     string         `json:"objName,omitempty" gorm:"size:200;comment:存储的对象名称"`
	CreateTime  int64          `json:"createTime" gorm:"autoCreateTime:milli"`
	DeleteTime  gorm.DeletedAt `json:"deleteTime,omitempty" gorm:"index"`
}

func (*AssetVersion) TableComment() string {
	return `资产文件版本表`
}

func (af *AssetVersion) UnmarshalJSON(b []byte) error {
	j := gjson.ParseBytes(b)
	af.ID = j.Get("id").Uint()
	af.FileID = j.Get("fileId").Uint()
	af.OssConfigID = j.Get("ossConfigId").Uint()
	af.ObjName = j.Get("objName").String()
	af.CreateTime = j.Get("createTime").Int()
	return nil
}

func init() {
	common.Models = append(common.Models, &AssetVersion{})
}
