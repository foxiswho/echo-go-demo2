package multitenant

type IMultiTenantPg interface {
}

// 多租户
type MultiTenantPg struct {
	//DomainId  int64 `json:"domainId"`  //领域id,区域id
	//TradeId  int64 `json:"tradeId"`  //运营商,交易中心,贸易商ID
	TenantId int64 `json:"tenantId"` //租户ID
	OwnerId  int64 `json:"ownerId"`  //货权方ID,卖家ID
}
