// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserProductConfig is the golang structure for table user_product_config.
type UserProductConfig struct {
	Id            uint        `json:"id"            ` //
	UserProductId int         `json:"userProductId" ` //
	Lport         string      `json:"lport"         ` // Local Port Begin 即终端代理要监听本地端口范围的开始值（包含开始值
	Rmt           string      `json:"rmt"           ` // remote是个数组，源站ip
	Status        int         `json:"status"        ` // 1是正常2是异常
	HostName      string      `json:"hostName"      ` // 1是服务器2是云
	HostType      int         `json:"hostType"      ` // 1是服务器2是云
	HostId        int         `json:"hostId"        ` // 产品id
	CreatedAt     *gtime.Time `json:"createdAt"     ` //
	UpdatedAt     *gtime.Time `json:"updatedAt"     ` //
	HostClass     int         `json:"hostClass"     ` // 1是服务器2是云
	CustomIp      string      `json:"customIp"      ` // 自定义ip
	LportInterval int         `json:"lportInterval" ` // 步阶
	PortLimit     int         `json:"portLimit"     ` // 端口限制个数
	GetIpWay      int         `json:"getIpWay"      ` // 获取真实ip方式：1不需要，2通过代理，3通过连接
	Type          int         `json:"type"          ` // 类型 0 PC端  1 移动端
	YdCustomIp    string      `json:"ydCustomIp"    ` // 移动端，客户连接ip改为127.0.0.1
}