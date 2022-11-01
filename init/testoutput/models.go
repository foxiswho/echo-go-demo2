package models

import (
	"time"
)

type IamAccount struct {
	Id               int64     `xorm:"pk BIGINT(20)"`
	CreateTime       time.Time `xorm:"default current_timestamp() comment('创建时间') index DATETIME"`
	UpdateTime       time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateId         int64     `xorm:"not null default 0 comment('创建人') index BIGINT(20)"`
	UpdateId         int64     `xorm:"not null default 0 comment('更新人') BIGINT(20)"`
	TenantId         int64     `xorm:"not null default 0 comment('租户') index BIGINT(20)"`
	OrgId            int64     `xorm:"not null default 0 comment('组织id') index BIGINT(20)"`
	Account          string    `xorm:"comment('账户') index CHAR(255)"`
	AccountMd5       string    `xorm:"comment('账户md5') CHAR(32)"`
	Password         string    `xorm:"comment('密码') CHAR(32)"`
	Salt             string    `xorm:"comment('组织id') CHAR(32)"`
	Phone            string    `xorm:"comment('手机号') index VARCHAR(50)"`
	Mail             string    `xorm:"comment('邮箱') index VARCHAR(255)"`
	MailMd5          string    `xorm:"CHAR(32)"`
	Code             string    `xorm:"comment('编码') CHAR(100)"`
	MailVerify       int       `xorm:"not null default 2 comment('邮箱验证1是2否') TINYINT(1)"`
	PhoneVerify      int       `xorm:"not null default 2 comment('手机验证1是2否') TINYINT(1)"`
	State            int       `xorm:"default 1 comment('启用1是2否') index TINYINT(1)"`
	RegisterTime     time.Time `xorm:"not null default '0001-01-01 00:00:00' comment('注册时间') DATETIME"`
	RegisterIp       string    `xorm:"comment('注册ip') CHAR(100)"`
	LoginTime        time.Time `xorm:"default '0001-01-01 00:00:00' comment('登陆时间') DATETIME"`
	RoleId           int64     `xorm:"not null default 0 comment('角色id') index BIGINT(20)"`
	DepartmentIdMain int64     `xorm:"not null default 0 comment('主部门id') index BIGINT(20)"`
	DepartmentIds    string    `xorm:"comment('部门') TEXT(65535)"`
	TeamIds          string    `xorm:"comment('团队') TEXT(65535)"`
	OrgIds           string    `xorm:"comment('团队') TEXT(65535)"`
	TypeDomain       int       `xorm:"not null default 0 comment('域类型') index INT(11)"`
}

type IamAccountLoginLog struct {
	Id          int64     `xorm:"pk BIGINT(20)"`
	CreateTime  time.Time `xorm:"default current_timestamp() comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"comment('更新时间') DATETIME"`
	CreateId    int64     `xorm:"not null default 0 comment('创建人') BIGINT(20)"`
	UpdateId    int64     `xorm:"not null default 0 comment('更新人') BIGINT(20)"`
	TenantId    int64     `xorm:"not null default 0 comment('租户') BIGINT(20)"`
	OrgId       int64     `xorm:"not null default 0 comment('组织id') BIGINT(20)"`
	LoginSource int64     `xorm:"not null default 0 comment('登陆来源') BIGINT(20)"`
}
