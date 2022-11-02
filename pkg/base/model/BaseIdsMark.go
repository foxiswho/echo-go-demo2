package model

import (
	"github.com/zxysilent/blog/pkg/base/holder/multitenant"
	"github.com/zxysilent/blog/pkg/base/holder/session"
)

// BaseIdMark 基础 详情
type BaseIdsMark[ID, MARK any] struct {
	Ids           []ID                       `json:"ids"`
	Mark          MARK                       `json:"mark"`
	MultiTenant   multitenant.IMultiTenantPg `json:"multiTenant"`   //多租户
	SessionHolder session.ISessionHolderPg   `json:"sessionHolder"` // 会话信息
}
