package model

import (
	"github.com/pangu-2/go-echo-demo/pkg/base/interfaces"
)

// BaseIds 基础 详情
type BaseIds[ID any] struct {
	Ids    []ID                 `json:"ids"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
}
