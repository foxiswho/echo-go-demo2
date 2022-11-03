package model

import (
	"encoding/json"

	"github.com/pangu-2/go-echo-demo/pkg/base/interfaces"
)

type BaseQueryByIdNotIdVdo[ID any] struct {
	Id     ID                   `json:"id"`
	IdNot  ID                   `json:"idNot"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
}

func (c BaseQueryByIdNotIdVdo) ToJsonString() (string, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
