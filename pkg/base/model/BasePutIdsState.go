package model

import (
	"github.com/zxysilent/blog/pkg/base/holder/multitenant"
	"github.com/zxysilent/blog/pkg/base/holder/session"
)

// BasePutIdsState 基础 状态更新
type BasePutIdsState[STATE, ID, D any] struct {
	State         STATE                      `json:"state"`
	Ids           []ID                       `json:"ids"`
	MultiTenant   multitenant.IMultiTenantPg `json:"multiTenant"`   //多租户
	SessionHolder session.ISessionHolderPg   `json:"sessionHolder"` // 会话信息
	Data          D                          `json:"data"`          //数据
	Msg           string                     `json:"msg"`           //消息
}
