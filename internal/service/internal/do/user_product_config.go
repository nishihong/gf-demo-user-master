// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserProductConfig is the golang structure of table yjs_user_product_config for DAO operations like Where/Data.
type UserProductConfig struct {
	g.Meta        `orm:"table:yjs_user_product_config, do:true"`
	Id            interface{} //
	UserProductId interface{} //
	Lport         interface{} // Local Port Begin 即终端代理要监听本地端口范围的开始值（包含开始值
	Rmt           interface{} // remote是个数组，源站ip
	Status        interface{} // 1是正常2是异常
	HostName      interface{} // 1是服务器2是云
	HostType      interface{} // 1是服务器2是云
	HostId        interface{} // 产品id
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
	HostClass     interface{} // 1是服务器2是云
	CustomIp      interface{} // 自定义ip
	LportInterval interface{} // 步阶
	PortLimit     interface{} // 端口限制个数
	GetIpWay      interface{} // 获取真实ip方式：1不需要，2通过代理，3通过连接
	Type          interface{} // 类型 0 PC端  1 移动端
	YdCustomIp    interface{} // 移动端，客户连接ip改为127.0.0.1
}