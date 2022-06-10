// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SdkUserProduct1 is the golang structure of table yjs_sdk_user_product1 for DAO operations like Where/Data.
type SdkUserProduct1 struct {
	g.Meta              `orm:"table:yjs_sdk_user_product1, do:true"`
	Id                  interface{} //
	UserId              interface{} // 用户id
	ProductId           interface{} // 产品id
	UserSerialId        interface{} // 用户序列号表id
	ProductType         interface{} // 对应产品表的类型
	HostName            interface{} // 1是服务器2是云
	HostType            interface{} // 1是服务器2是云
	HostId              interface{} // 产品id
	OnOff               interface{} // -1是未初始化，1是启用，2是关闭
	Status              interface{} // 用户产品状态1正常 2关闭 -1未支付
	Name                interface{} // 产品名称
	AvailablePoints     interface{} // 可用点数
	AvaliableFrequency  interface{} // 可用封装次数
	CreatedAt           *gtime.Time //
	UpdatedAt           *gtime.Time //
	StartAt             *gtime.Time // 生效日期
	EndAt               *gtime.Time // 失效日期
	Bz                  interface{} //
	ClientLimit         interface{} // 客户端最大连接数
	SharedNodesNum      interface{} // 共享节点数
	IndependentNodesNum interface{} // 独立节点数
	CountUpLink         interface{} // 统计数据
	CountDownLink       interface{} // 统计数据
	CountConnections    interface{} // 统计数据
	CountAt             *gtime.Time // 统计数据时间
	IsLogo              interface{} // 是否展示logo,1展示，0不展示
	Username            interface{} // sdk字段-所属用户
	ProductNum          interface{} // sdk字段-产品编号
	Amount              interface{} // sdk字段-金额
	TradedMode          interface{} // sdk字段-交易方式
	TradedType          interface{} // sdk字段-交易类型
	OrderType           interface{} // sdk字段-订单类型
	DeletedAt           *gtime.Time //
	Operator            interface{} // SDK字段
	SourceType          interface{} // 客户来源 1云加速 2sdk
	SdkNum              interface{} // sdk_num
	Guid                interface{} // 规则改变标识
	SkinFile            interface{} // 客户端皮肤资源包 文件ID
	IcoImg              interface{} // 托盘图标 文件ID
	Remarks             interface{} // 数据中心-列表备注
	RealAmount          interface{} // 数据中心-实际支付价格
	SellId              interface{} // 销售id
	NodeLimit           interface{} // 节点下发数量限制
	GameNum             interface{} // 游戏数量限制
}
