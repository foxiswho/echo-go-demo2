package model

import "encoding/json"

type BaseQueryVdo struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
	Deleted  int `json:"deleted"`
	Opened   int `json:"opened"`
}

func (c BaseQueryVdo) ToJsonString() (string, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
