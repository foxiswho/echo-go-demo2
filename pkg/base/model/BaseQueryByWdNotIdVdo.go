package model

import (
	"encoding/json"

	"github.com/pangu-2/go-echo-demo/pkg/base/interfaces"
)

type BaseQueryByWdNotIdVdo[ID any] struct {
	Wd     string               `json:"wd"`
	IdNot  ID                   `json:"idNot"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
}

func (c BaseQueryByWdNotIdVdo) ToJsonString() (string, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
