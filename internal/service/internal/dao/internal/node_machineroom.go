// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NodeMachineroomDao is the data access object for table yjs_node_machineroom.
type NodeMachineroomDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns NodeMachineroomColumns // columns contains all the column names of Table for convenient usage.
}

// NodeMachineroomColumns defines and stores column names for table yjs_node_machineroom.
type NodeMachineroomColumns struct {
	Id            string //
	NodeId        string // 节点id
	MachineroomId string // 序列号id
	CreatedAt     string //
	UpdatedAt     string //
}

//  nodeMachineroomColumns holds the columns for table yjs_node_machineroom.
var nodeMachineroomColumns = NodeMachineroomColumns{
	Id:            "id",
	NodeId:        "node_id",
	MachineroomId: "machineroom_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewNodeMachineroomDao creates and returns a new DAO object for table data access.
func NewNodeMachineroomDao() *NodeMachineroomDao {
	return &NodeMachineroomDao{
		group:   "default",
		table:   "yjs_node_machineroom",
		columns: nodeMachineroomColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NodeMachineroomDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NodeMachineroomDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NodeMachineroomDao) Columns() NodeMachineroomColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NodeMachineroomDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NodeMachineroomDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NodeMachineroomDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}