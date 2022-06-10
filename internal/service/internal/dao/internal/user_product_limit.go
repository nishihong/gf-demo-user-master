// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserProductLimitDao is the data access object for table yjs_user_product_limit.
type UserProductLimitDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns UserProductLimitColumns // columns contains all the column names of Table for convenient usage.
}

// UserProductLimitColumns defines and stores column names for table yjs_user_product_limit.
type UserProductLimitColumns struct {
	Id            string //
	UserProductId string // 产品id
	ServerId      string // 服务器id
	IsOpen        string // 状态
	RunningTime   string // 运行时间
	MaxConnect    string // 一定时间内最大连接数
	ShieldTime    string // 屏蔽时间
	CreatedAt     string //
	UpdatedAt     string //
}

//  userProductLimitColumns holds the columns for table yjs_user_product_limit.
var userProductLimitColumns = UserProductLimitColumns{
	Id:            "id",
	UserProductId: "user_product_id",
	ServerId:      "server_id",
	IsOpen:        "is_open",
	RunningTime:   "running_time",
	MaxConnect:    "max_connect",
	ShieldTime:    "shield_time",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewUserProductLimitDao creates and returns a new DAO object for table data access.
func NewUserProductLimitDao() *UserProductLimitDao {
	return &UserProductLimitDao{
		group:   "default",
		table:   "yjs_user_product_limit",
		columns: userProductLimitColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserProductLimitDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserProductLimitDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserProductLimitDao) Columns() UserProductLimitColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserProductLimitDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserProductLimitDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserProductLimitDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}