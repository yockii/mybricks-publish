package common

import "manatee-publish/pkg/server"

type BaseDomain[T Model] interface {
	GetModel() T
	GetOrderBy() string
	GetTimeConditionList() map[string]*server.TimeCondition
}
