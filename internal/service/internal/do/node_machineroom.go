// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NodeMachineroom is the golang structure of table yjs_node_machineroom for DAO operations like Where/Data.
type NodeMachineroom struct {
	g.Meta        `orm:"table:yjs_node_machineroom, do:true"`
	Id            interface{} //
	NodeId        interface{} // 节点id
	MachineroomId interface{} // 序列号id
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
