package model

import (
	"encoding/json"

	"github.com/pangu-2/go-echo-demo/pkg/base/holder"
	"github.com/pangu-2/go-echo-demo/pkg/base/holder/multitenant"
)

type BaseQueryByWdNotIdVdo[ID any] struct {
	Wd            string                     `json:"wd"`
	IdNot         ID                         `json:"idNot"`
	MultiTenant   multitenant.IMultiTenantPg `json:"multiTenant"`   //多租户
	SessionHolder holder.IHolderPg           `json:"sessionHolder"` // 会话信息
}

func (c BaseQueryByWdNotIdVdo) ToJsonString() (string, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
