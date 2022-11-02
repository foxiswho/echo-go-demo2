package upload

type UploadSuccess struct {
	No string `json:"no"`
}

func NewUploadSuccess(no string) *UploadSuccess {
	n := new(UploadSuccess)
	n.No = no
	return n
}

func (c UploadSuccess) Get() UploadSuccess {
	return c
}
