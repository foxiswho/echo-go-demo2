package model

import (
	"encoding/json"

	"github.com/zxysilent/blog/pkg/base/holder/multitenant"
	"github.com/zxysilent/blog/pkg/base/holder/session"
)

type BaseQueryByIdThatIdVdo[ID any] struct {
	Id            ID                         `json:"id"`
	IdThat        ID                         `json:"idThat"`
	MultiTenant   multitenant.IMultiTenantPg `json:"multiTenant"`   //多租户
	SessionHolder session.ISessionHolderPg   `json:"sessionHolder"` // 会话信息
}

func (c BaseQueryByIdThatIdVdo) ToJsonString() (string, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
