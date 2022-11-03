package interfaces

// 实体总
type IHolderPg interface {
}

// 多租户
type IMultiTenantPg interface {
}

// 用户数据
type IHolderDataPg interface {
}

// jwt数据
type IHolderJwtPg interface {
}

// 租户数据
type ITenantPg interface {
}

// 其他
type IHolderOtherPg interface {
}

type StandardHolder struct {
	Jwt         *IHolderJwtPg          `json:"jwt,omitempty" commont:"jwt"`
	MultiTenant *IMultiTenantPg        `json:"mult,omitempty" commont:"多租户"`
	HolderData  *IHolderDataPg         `json:"hdata,omitempty" commont:"用户数据"`
	Tenant      *ITenantPg             `json:"tenant,omitempty" commont:"租户数据"`
	Other       *IHolderOtherPg        `json:"other,omitempty" commont:"其他数据"`
	Map         map[string]interface{} `json:"map,omitempty" commont:"map"`
}
