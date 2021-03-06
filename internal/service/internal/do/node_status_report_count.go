// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NodeStatusReportCount is the golang structure of table yjs_node_status_report_count for DAO operations like Where/Data.
type NodeStatusReportCount struct {
	g.Meta      `orm:"table:yjs_node_status_report_count, do:true"`
	Id          interface{} //
	NodeId      interface{} // 节点id
	SourceType  interface{} // 数据来源类型，1云加速节点，2游戏盾节点
	Connections interface{} // 并发数
	BytesSent   interface{} // 发送字节数
	BytesRecv   interface{} // 接收字节数
	PacketsSent interface{} // 发送数据包数量
	PacketsRecv interface{} // 接收数据包数量
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
