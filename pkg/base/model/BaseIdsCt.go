package model

// BaseIdCt 基础 详情
type BaseIdsCt[ID any] struct {
	Ids []ID `json:"ids"`
}
