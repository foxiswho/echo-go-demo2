package model

import (
	"github.com/pangu-2/go-echo-demo/pkg/base/holder"
	"github.com/pangu-2/go-echo-demo/pkg/base/holder/multitenant"
)

// BaseIdMarks 基础 详情
type BaseIdsMarks[ID, MARK any] struct {
	Ids           []ID                       `json:"ids"`
	Marks         []MARK                     `json:"marks"`
	MultiTenant   multitenant.IMultiTenantPg `json:"multiTenant"`   //多租户
	SessionHolder holder.IHolderPg           `json:"sessionHolder"` // 会话信息
}
