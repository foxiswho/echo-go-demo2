package model

import (
	"github.com/zxysilent/blog/pkg/base/holder/multitenant"
	"github.com/zxysilent/blog/pkg/base/holder/session"
)

// BaseIdMarks 基础 详情
type BaseIdsMarks[ID, MARK any] struct {
	Ids           []ID                       `json:"ids"`
	Marks         []MARK                     `json:"marks"`
	MultiTenant   multitenant.IMultiTenantPg `json:"multiTenant"`   //多租户
	SessionHolder session.ISessionHolderPg   `json:"sessionHolder"` // 会话信息
}
