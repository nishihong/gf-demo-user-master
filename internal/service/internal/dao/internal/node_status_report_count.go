// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NodeStatusReportCountDao is the data access object for table yjs_node_status_report_count.
type NodeStatusReportCountDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns NodeStatusReportCountColumns // columns contains all the column names of Table for convenient usage.
}

// NodeStatusReportCountColumns defines and stores column names for table yjs_node_status_report_count.
type NodeStatusReportCountColumns struct {
	Id          string //
	NodeId      string // 节点id
	SourceType  string // 数据来源类型，1云加速节点，2游戏盾节点
	Connections string // 并发数
	BytesSent   string // 发送字节数
	BytesRecv   string // 接收字节数
	PacketsSent string // 发送数据包数量
	PacketsRecv string // 接收数据包数量
	CreatedAt   string //
	UpdatedAt   string //
}

//  nodeStatusReportCountColumns holds the columns for table yjs_node_status_report_count.
var nodeStatusReportCountColumns = NodeStatusReportCountColumns{
	Id:          "id",
	NodeId:      "node_id",
	SourceType:  "source_type",
	Connections: "connections",
	BytesSent:   "bytes_sent",
	BytesRecv:   "bytes_recv",
	PacketsSent: "packets_sent",
	PacketsRecv: "packets_recv",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewNodeStatusReportCountDao creates and returns a new DAO object for table data access.
func NewNodeStatusReportCountDao() *NodeStatusReportCountDao {
	return &NodeStatusReportCountDao{
		group:   "default",
		table:   "yjs_node_status_report_count",
		columns: nodeStatusReportCountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NodeStatusReportCountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NodeStatusReportCountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NodeStatusReportCountDao) Columns() NodeStatusReportCountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NodeStatusReportCountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NodeStatusReportCountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NodeStatusReportCountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
