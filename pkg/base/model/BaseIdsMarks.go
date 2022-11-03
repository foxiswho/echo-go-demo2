package model

import (
	"github.com/pangu-2/go-echo-demo/pkg/base/interfaces"
)

// BaseIdMarks 基础 详情
type BaseIdsMarks[ID, MARK any] struct {
	Ids    []ID                 `json:"ids"`
	Marks  []MARK               `json:"marks"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
}
