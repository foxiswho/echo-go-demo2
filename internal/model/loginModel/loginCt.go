package loginModel

type LoginCt struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
	Code     string `json:"code" `
	CodeSms  string `json:"codeSms" `
	Type     string `json:"type" `
}
