package jwtHolder

// jwt 多租户
type JwtMultiTenantPg struct {
	TenantId int64 `json:"tId"` //租户ID
	OrgId    int64 `json:"oId"` //组织
}
