package model

import (
	"github.com/pangu-2/go-echo-demo/pkg/base/interfaces"
)

// BasePutIdsState 基础 状态更新
type BasePutIdsState[STATE, ID, D any] struct {
	State  STATE                `json:"state"`
	Ids    []ID                 `json:"ids"`
	Holder interfaces.IHolderPg `json:"holder"` // 会话信息
	Data   D                    `json:"data"`   //数据
	Msg    string               `json:"msg"`    //消息
}
