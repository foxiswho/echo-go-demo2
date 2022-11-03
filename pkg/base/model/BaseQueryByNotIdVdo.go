package model

import (
	"encoding/json"

	"github.com/pangu-2/go-echo-demo/pkg/base/holder"
	"github.com/pangu-2/go-echo-demo/pkg/base/holder/multitenant"
)

type BaseQueryByNotIdVdo[ID any] struct {
	IdNot         ID                         `json:"idNot"`
	MultiTenant   multitenant.IMultiTenantPg `json:"multiTenant"`   //多租户
	SessionHolder holder.IHolderPg           `json:"sessionHolder"` // 会话信息
}

func (c BaseQueryByNotIdVdo) ToJsonString() (string, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
