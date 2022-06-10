// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SdkUserProduct is the golang structure of table yjs_sdk_user_product for DAO operations like Where/Data.
type SdkUserProduct struct {
	g.Meta      `orm:"table:yjs_sdk_user_product, do:true"`
	Id          interface{} //
	UserId      interface{} // 用户id
	DeviceId    interface{} // 实例id
	Lip         interface{} // 客户连接ip
	ProductId   interface{} // 产品id
	OnOff       interface{} // -1是未初始化，1是启用，2是关闭
	Status      interface{} // 用户产品状态1正常 2关闭 -1未支付
	Name        interface{} // 产品名称
	StartAt     *gtime.Time // 生效日期
	EndAt       *gtime.Time // 失效日期
	Bz          interface{} // 备注
	ClientLimit interface{} // 客户端最大连接数
	Username    interface{} // sdk字段-所属用户
	Amount      interface{} // sdk字段-金额
	TradedMode  interface{} // sdk字段-交易方式
	TradedType  interface{} // sdk字段-交易类型
	OrderType   interface{} // sdk字段-订单类型
	Operator    interface{} // SDK字段-操作员
	SourceType  interface{} // 客户来源 1云加速 2sdk
	Remarks     interface{} // 数据中心-列表备注
	RealAmount  interface{} // 数据中心-实际支付价格
	SellId      interface{} // 销售id
	NodeLimit   interface{} // 节点下发数量限制
	DayRemark   interface{} // 日常备注（数据中心使用）
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
}