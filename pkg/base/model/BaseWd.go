package model

import (
	"github.com/pangu-2/go-echo-demo/pkg/base/interfaces"
)

// BaseWd 基础 详情
type BaseWd struct {
	Wd     string               `json:"wd"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
}
