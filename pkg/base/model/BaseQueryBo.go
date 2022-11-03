package model

import (
	"encoding/json"

	"github.com/pangu-2/go-echo-demo/pkg/base/interfaces"
)

type BaseQueryBo struct {
	pageNum  int                  `json:"pageNum"`
	PageSize int                  `json:"pageSize"`
	Holder   interfaces.IHolderPg `json:"holder"` // 会话信息
}

func (c BaseQueryBo) ToJsonString() (string, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
