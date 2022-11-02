package model

import (
	"github.com/zxysilent/blog/pkg/base/holder/multitenant"
	"github.com/zxysilent/blog/pkg/base/holder/session"
)

// BaseWdArr 基础 详情
type BaseWdArr struct {
	Wd            []string                   `json:"wd"`
	MultiTenant   multitenant.IMultiTenantPg `json:"multiTenant"`   //多租户
	SessionHolder session.ISessionHolderPg   `json:"sessionHolder"` // 会话信息
}
