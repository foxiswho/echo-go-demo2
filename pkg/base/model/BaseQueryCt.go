package model

import "encoding/json"

type BaseQueryCt struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

func (c BaseQueryCt) ToJsonString() (string, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
