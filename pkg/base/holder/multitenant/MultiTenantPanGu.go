package multitenant

import "github.com/pangu-2/go-echo-demo/pkg/base/interfaces"

type IMultiTenantPg interface {
	*interfaces.IMultiTenantPg
}

// 多租户
type MultiTenantPg struct {
	//DomainId  int64 `json:"domainId"`  //领域id,区域id
	//TradeId  int64 `json:"tradeId"`  //运营商,交易中心,贸易商ID
	// OwnerId  int64 `json:"ownerId"`  //货权方ID,卖家ID
	TenantId int64 `json:"tenantId,omitempty"` //租户ID
	OrgId    int64 `json:"oId,omitempty"`      //组织
}
