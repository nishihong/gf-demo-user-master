// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NodesProductDao is the data access object for table yjs_nodes_product.
type NodesProductDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns NodesProductColumns // columns contains all the column names of Table for convenient usage.
}

// NodesProductColumns defines and stores column names for table yjs_nodes_product.
type NodesProductColumns struct {
	Id        string //
	NodeId    string // 节点id
	ProductId string // 套餐类型  'product_type'    => [          '1'=>"基础版",          '2'=>"稳定版",          '3'=>"专业版",          '4'=>"高级版",          '5'=>"定制版",      ],
	CreatedAt string //
	UpdatedAt string //
}

//  nodesProductColumns holds the columns for table yjs_nodes_product.
var nodesProductColumns = NodesProductColumns{
	Id:        "id",
	NodeId:    "node_id",
	ProductId: "product_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewNodesProductDao creates and returns a new DAO object for table data access.
func NewNodesProductDao() *NodesProductDao {
	return &NodesProductDao{
		group:   "default",
		table:   "yjs_nodes_product",
		columns: nodesProductColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NodesProductDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NodesProductDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NodesProductDao) Columns() NodesProductColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NodesProductDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NodesProductDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NodesProductDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}