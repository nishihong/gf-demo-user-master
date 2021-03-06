// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NodesProductTypeDao is the data access object for table yjs_nodes_product_type.
type NodesProductTypeDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns NodesProductTypeColumns // columns contains all the column names of Table for convenient usage.
}

// NodesProductTypeColumns defines and stores column names for table yjs_nodes_product_type.
type NodesProductTypeColumns struct {
	Id            string //
	NodeId        string // 节点id
	ProductTypeId string // 套餐类型  'product_type'    => [          '1'=>"基础版",          '2'=>"稳定版",          '3'=>"专业版",          '4'=>"高级版",          '5'=>"定制版",      ],
	CreatedAt     string //
	UpdatedAt     string //
}

//  nodesProductTypeColumns holds the columns for table yjs_nodes_product_type.
var nodesProductTypeColumns = NodesProductTypeColumns{
	Id:            "id",
	NodeId:        "node_id",
	ProductTypeId: "product_type_id",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewNodesProductTypeDao creates and returns a new DAO object for table data access.
func NewNodesProductTypeDao() *NodesProductTypeDao {
	return &NodesProductTypeDao{
		group:   "default",
		table:   "yjs_nodes_product_type",
		columns: nodesProductTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NodesProductTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NodesProductTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NodesProductTypeDao) Columns() NodesProductTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NodesProductTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NodesProductTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NodesProductTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
