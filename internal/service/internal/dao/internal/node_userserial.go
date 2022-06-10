// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NodeUserserialDao is the data access object for table yjs_node_userserial.
type NodeUserserialDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns NodeUserserialColumns // columns contains all the column names of Table for convenient usage.
}

// NodeUserserialColumns defines and stores column names for table yjs_node_userserial.
type NodeUserserialColumns struct {
	Id               string //
	NodeId           string // 节点id
	UserSerialId     string // 序列号id
	CreatedAt        string //
	UpdatedAt        string //
	MachineId        string // 机房id
	DistributionType string // 1手动分配，2自动分配
	Type             string // 1共享。2 独立
	NodeType         string // 1防御网关。2 安全网关
}

//  nodeUserserialColumns holds the columns for table yjs_node_userserial.
var nodeUserserialColumns = NodeUserserialColumns{
	Id:               "id",
	NodeId:           "node_id",
	UserSerialId:     "user_serial_id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	MachineId:        "machine_id",
	DistributionType: "distribution_type",
	Type:             "type",
	NodeType:         "node_type",
}

// NewNodeUserserialDao creates and returns a new DAO object for table data access.
func NewNodeUserserialDao() *NodeUserserialDao {
	return &NodeUserserialDao{
		group:   "default",
		table:   "yjs_node_userserial",
		columns: nodeUserserialColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NodeUserserialDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NodeUserserialDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NodeUserserialDao) Columns() NodeUserserialColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NodeUserserialDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NodeUserserialDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NodeUserserialDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
