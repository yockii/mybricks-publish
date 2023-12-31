package common

import "manatee-publish/pkg/util"

type Model interface {
	TableComment() string
	AddRequired() string
	InstanceRequired() string
	CheckDuplicatedModel() Model
	UpdateConditionModel() Model
	UpdateModel() Model
	InitDefaultFields()
	UpdateRequired() string
	FuzzyQueryMap() map[string]string
	ExactMatchModel() Model
	DeleteRequired() string
	ListOmits() string
}

var Models []Model

type BaseModel struct {
	ID uint64 `json:"id,omitempty,string" gorm:"primaryKey;autoIncrement:false"`
}

func (*BaseModel) TableComment() string {
	return "空表"
}
func (m *BaseModel) AddRequired() string {
	return ""
}
func (m *BaseModel) InstanceRequired() string {
	if m.ID == 0 {
		return "id"
	}
	return ""
}
func (m *BaseModel) CheckDuplicatedModel() Model {
	return nil
}
func (m *BaseModel) UpdateConditionModel() Model {
	b := new(BaseModel)
	b.ID = m.ID
	return b
}
func (m *BaseModel) UpdateModel() Model {
	panic("implement me")
}
func (m *BaseModel) InitDefaultFields() {
	m.ID = util.SnowflakeId()
}
func (m *BaseModel) UpdateRequired() string {
	if m.ID == 0 {
		return "id"
	}
	return ""
}
func (m *BaseModel) DeleteRequired() string {
	if m.ID == 0 {
		return "id"
	}
	return ""
}
func (m *BaseModel) FuzzyQueryMap() map[string]string {
	return nil
}
func (m *BaseModel) ExactMatchModel() Model {
	return m
}
func (m *BaseModel) ListOmits() string {
	return ""
}
