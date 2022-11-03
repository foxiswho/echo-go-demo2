package model

import (
	"github.com/pangu-2/go-echo-demo/pkg/base/holder"
	"github.com/pangu-2/go-echo-demo/pkg/base/holder/multitenant"
)

// BaseId 基础 详情
type BaseId[ID any] struct {
	Id            ID                         `json:"id"`
	MultiTenant   multitenant.IMultiTenantPg `json:"multiTenant"`   //多租户
	SessionHolder holder.IHolderPg           `json:"sessionHolder"` // 会话信息
}
