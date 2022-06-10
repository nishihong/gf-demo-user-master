// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NodesUserSerial is the golang structure of table yjs_nodes_user_serial for DAO operations like Where/Data.
type NodesUserSerial struct {
	g.Meta           `orm:"table:yjs_nodes_user_serial, do:true"`
	Id               interface{} //
	NodeId           interface{} // 节点id
	MachineId        interface{} // 机房id
	UserSerialId     interface{} // 用户产品id
	DistributionType interface{} // 1手动分配，2自动分配
	Type             interface{} // 1共享。2 独立
	NodeType         interface{} // 1防御网关。2 安全网关
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
}