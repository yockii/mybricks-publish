package model

import (
	"github.com/tidwall/gjson"
	"manatee-publish/pkg/common"
	"strconv"
)

const (
	UserStatusNormal   = 1
	UserStatusDisabled = -1
)

type User struct {
	common.BaseModel
	Username   string `json:"username,omitempty" gorm:"size:30;index;comment:用户名"`
	Password   string `json:"password,omitempty" gorm:"size:100;comment:密码"`
	NickName   string `json:"nickName,omitempty" gorm:"size:50;comment:昵称"`
	Email      string `json:"email,omitempty" gorm:"size:50;comment:邮箱"`
	Mobile     string `json:"mobile,omitempty" gorm:"size:50;comment:手机号"`
	Status     int    `json:"status,omitempty" gorm:"comment:状态 1-正常"`
	CreateTime int64  `json:"createTime" gorm:"autoCreateTime:milli"`
	UpdateTime int64  `json:"updateTime" gorm:"autoUpdateTime:milli"`
}

func (_ *User) TableComment() string {
	return "用户表"
}

func (u *User) UnmarshalJSON(b []byte) error {
	j := gjson.ParseBytes(b)
	u.ID = j.Get("id").Uint()
	u.Username = j.Get("username").String()
	u.Password = j.Get("password").String()
	u.NickName = j.Get("nickName").String()
	u.Email = j.Get("email").String()
	u.Mobile = j.Get("mobile").String()
	u.Status = int(j.Get("status").Int())
	return nil
}

func (u *User) AddRequired() string {
	if u.Username == "" || u.Password == "" {
		return "username, password"
	}
	return ""
}
func (u *User) CheckDuplicatedModel() common.Model {
	return &User{
		Username: u.Username,
	}
}
func (m *User) InitDefaultFields() {
	m.BaseModel.InitDefaultFields()
	if m.NickName == "" {
		m.NickName = "用户" + strconv.FormatUint(m.ID, 10)
	}
}
func (u *User) UpdateModel() common.Model {
	return &User{
		NickName: u.NickName,
		Email:    u.Email,
		Mobile:   u.Mobile,
		Status:   u.Status,
	}
}
func (u *User) FuzzyQueryMap() map[string]string {
	result := make(map[string]string)
	if u.Username != "" {
		result["username"] = "%" + u.Username + "%"
	}
	if u.NickName != "" {
		result["nick_name"] = "%" + u.NickName + "%"
	}
	if u.Email != "" {
		result["email"] = "%" + u.Email + "%"
	}
	if u.Mobile != "" {
		result["mobile"] = "%" + u.Mobile + "%"
	}
	return result
}
func (u *User) ExactMatchModel() common.Model {
	return &User{
		Status: u.Status,
	}
}
func (m *User) ListOmits() string {
	return "password"
}

func init() {
	common.Models = append(common.Models, &User{})
}
