package model

import (
	"github.com/pangu-2/go-echo-demo/pkg/base/holder"
	"github.com/pangu-2/go-echo-demo/pkg/base/holder/multitenant"
)

// BaseIdMark 基础 详情
type BaseIdsMark[ID, MARK any] struct {
	Ids           []ID                       `json:"ids"`
	Mark          MARK                       `json:"mark"`
	MultiTenant   multitenant.IMultiTenantPg `json:"multiTenant"`   //多租户
	SessionHolder holder.IHolderPg           `json:"sessionHolder"` // 会话信息
}
