// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAccount is the golang structure for table user_account.
type UserAccount struct {
	Id                int         `json:"id"                ` // 账单ID
	UserId            int         `json:"userId"            ` // 用户ID
	UserProductId     int         `json:"userProductId"     ` // 用户产品id
	ProductId         int         `json:"productId"         ` // 产品id
	AccountStyle      int         `json:"accountStyle"      ` // 操作方式 1进账 2出账
	AccountType       int         `json:"accountType"       ` // 1.用户充值，2.产品购买，3.产品续费
	AccountMoney      float64     `json:"accountMoney"      ` // 操作金额
	AccountIncometime int         `json:"accountIncometime" ` // 入款时间
	AccountStatus     int         `json:"accountStatus"     ` // 审核状态 -1待审核 1通过 2不通过 3审核中
	Content           string      `json:"content"           ` // 备注说明
	Oldmoney          float64     `json:"oldmoney"          ` // 用户操作之前的余额
	Nowmoney          float64     `json:"nowmoney"          ` // 用户操作之后的余额
	NumberType        int         `json:"numberType"        ` // 交易流水号的类型 （支付宝，微信，银行，其他等）
	Number            string      `json:"number"            ` // 交易流水号
	CashNumber        string      `json:"cashNumber"        ` // 代金券编码
	CashMoney         float64     `json:"cashMoney"         ` // 代金券使用金额
	CreatedAt         *gtime.Time `json:"createdAt"         ` //
	UpdatedAt         *gtime.Time `json:"updatedAt"         ` //
	Status            string      `json:"status"            ` // 入账状态 -1是未开始， 1是成功，2是失败
	Input             string      `json:"input"             ` // 详情
	Month             int         `json:"month"             ` // 续费月份
	UpgradeProductId  int         `json:"upgradeProductId"  ` // 升级产品id
	OldData           string      `json:"oldData"           ` // 旧数据
	IsAuto            int         `json:"isAuto"            ` // 是否自动生成 0->否 1->是
	StrategyMoney     float64     `json:"strategyMoney"     ` // 折扣
}
