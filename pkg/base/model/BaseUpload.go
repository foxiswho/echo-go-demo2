package model

import (
	"encoding/json"
	"mime/multipart"

	"github.com/zxysilent/blog/pkg/base/holder/session"
)

type BaseUploadBo struct {
	File          *multipart.FileHeader
	SessionHolder session.ISessionHolderPg `json:"sessionHolder"` // 会话信息
}
type Option func(*BaseUploadBo)

func NewBaseUpload(opts ...Option) *BaseUploadBo {
	s := new(BaseUploadBo)
	if len(opts) > 0 {
		for _, o := range opts {
			o(s)
		}
	}
	return s
}

func (c BaseUploadBo) ToJsonString() (string, error) {
	marshal, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
