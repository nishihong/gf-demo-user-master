// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserProductDownDao is the data access object for table yjs_user_product_down.
type UserProductDownDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns UserProductDownColumns // columns contains all the column names of Table for convenient usage.
}

// UserProductDownColumns defines and stores column names for table yjs_user_product_down.
type UserProductDownColumns struct {
	Id          string //
	ProductType string // 原版本
	DowngradeId string // 降级版本
	Serial      string // 序列号
	UserName    string // 用户名
	UserId      string // 用户id
	SellId      string // 销售id
	DownAt      string // 降级时间
	StartAt     string // 开始时间
	EndAt       string // 到期时间
	CreatedAt   string // 创建时间
	UpdatedAt   string //
	DeletedAt   string //
}

//  userProductDownColumns holds the columns for table yjs_user_product_down.
var userProductDownColumns = UserProductDownColumns{
	Id:          "id",
	ProductType: "product_type",
	DowngradeId: "downgrade_id",
	Serial:      "serial",
	UserName:    "user_name",
	UserId:      "user_id",
	SellId:      "sell_id",
	DownAt:      "down_at",
	StartAt:     "start_at",
	EndAt:       "end_at",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewUserProductDownDao creates and returns a new DAO object for table data access.
func NewUserProductDownDao() *UserProductDownDao {
	return &UserProductDownDao{
		group:   "default",
		table:   "yjs_user_product_down",
		columns: userProductDownColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserProductDownDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserProductDownDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserProductDownDao) Columns() UserProductDownColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserProductDownDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserProductDownDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserProductDownDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
