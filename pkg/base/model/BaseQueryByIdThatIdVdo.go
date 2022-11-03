package model

import (
	"encoding/json"

	"github.com/pangu-2/go-echo-demo/pkg/base/interfaces"
)

type BaseQueryByIdThatIdVdo[ID any] struct {
	Id     ID                   `json:"id"`
	IdThat ID                   `json:"idThat"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
}

func (c BaseQueryByIdThatIdVdo) ToJsonString() (string, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
