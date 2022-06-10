// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserProductCountDao is the data access object for table yjs_user_product_count.
type UserProductCountDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns UserProductCountColumns // columns contains all the column names of Table for convenient usage.
}

// UserProductCountColumns defines and stores column names for table yjs_user_product_count.
type UserProductCountColumns struct {
	Id               string //
	UserProductCount string //
	YxdCount         string //
	YjsCount         string //
	CreatedAt        string //
	UpdatedAt        string //
}

//  userProductCountColumns holds the columns for table yjs_user_product_count.
var userProductCountColumns = UserProductCountColumns{
	Id:               "id",
	UserProductCount: "user_product_count",
	YxdCount:         "yxd_count",
	YjsCount:         "yjs_count",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewUserProductCountDao creates and returns a new DAO object for table data access.
func NewUserProductCountDao() *UserProductCountDao {
	return &UserProductCountDao{
		group:   "default",
		table:   "yjs_user_product_count",
		columns: userProductCountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserProductCountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserProductCountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserProductCountDao) Columns() UserProductCountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserProductCountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserProductCountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserProductCountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
