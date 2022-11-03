package jwtHolder

// JwtPg 用户 会话信息 登录人信息
type JwtPg struct {
	MultiTenant   JwtMultiTenantPg `json:"mTenant" multiTenant` //多租户
	LoginNo       string           `json:"loginNo"`             //登录用户No,随时可以修改变动
	No            string           `json:"no"`                  //系统编号
	LoginUserName string           `json:"loginUserName"`       //登录用户名
	Name          string           `json:"name"`                //显示名称
	OrgName       string           `json:"OrgName"`             //组织名称
	TenantName    string           `json:"tName"`               //组织名称
	Type          int64            `json:"type"`                //类型
	Other         string           `json:"other"`               //其他信息
	Version       string           `json:"version"`             //版本
	Extra         interface{}      `json:"extra"`               //额外的，扩展
}

func NewJwtPg() *JwtPg {
	return new(JwtPg)
}
