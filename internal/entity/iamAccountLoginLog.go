package entity

import "time"

type IamAccountLoginLog struct {
	Id          int64     `xorm:"pk BIGINT(20)"`
	CreateTime  time.Time `xorm:"created default current_timestamp() comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateId    int64     `xorm:"not null default 0 comment('创建人') BIGINT(20)"`
	UpdateId    int64     `xorm:"not null default 0 comment('更新人') BIGINT(20)"`
	TenantId    int64     `xorm:"not null default 0 comment('租户') BIGINT(20)"`
	OrgId       int64     `xorm:"not null default 0 comment('组织id') BIGINT(20)"`
	LoginSource int64     `xorm:"not null default 0 comment('登陆来源') BIGINT(20)"`
}

func (IamAccountLoginLog) TableName() string {
	return "iam_account_login_log"
}
