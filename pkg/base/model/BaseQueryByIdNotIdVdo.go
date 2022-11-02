package model

import (
	"encoding/json"

	"github.com/zxysilent/blog/pkg/base/holder/multitenant"
	"github.com/zxysilent/blog/pkg/base/holder/session"
)

type BaseQueryByIdNotIdVdo[ID any] struct {
	Id            ID                         `json:"id"`
	IdNot         ID                         `json:"idNot"`
	MultiTenant   multitenant.IMultiTenantPg `json:"multiTenant"`   //多租户
	SessionHolder session.ISessionHolderPg   `json:"sessionHolder"` // 会话信息
}

func (c BaseQueryByIdNotIdVdo) ToJsonString() (string, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
