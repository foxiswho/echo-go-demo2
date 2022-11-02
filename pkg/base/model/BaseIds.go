package model

import (
	"github.com/zxysilent/blog/pkg/base/holder/multitenant"
	"github.com/zxysilent/blog/pkg/base/holder/session"
)

// BaseIds 基础 详情
type BaseIds[ID any] struct {
	Ids           []ID                       `json:"ids"`
	MultiTenant   multitenant.IMultiTenantPg `json:"multiTenant"`   //多租户
	SessionHolder session.ISessionHolderPg   `json:"sessionHolder"` // 会话信息
}
