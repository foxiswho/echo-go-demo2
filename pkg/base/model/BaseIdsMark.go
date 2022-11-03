package model

import (
	"github.com/pangu-2/go-echo-demo/pkg/base/interfaces"
)

// BaseIdMark 基础 详情
type BaseIdsMark[ID, MARK any] struct {
	Ids    []ID                 `json:"ids"`
	Mark   MARK                 `json:"mark"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
}
