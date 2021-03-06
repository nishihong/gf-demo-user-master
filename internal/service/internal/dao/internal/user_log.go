// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLogDao is the data access object for table yjs_user_log.
type UserLogDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns UserLogColumns // columns contains all the column names of Table for convenient usage.
}

// UserLogColumns defines and stores column names for table yjs_user_log.
type UserLogColumns struct {
	Id          string //
	UserName    string // 用户名
	Ip          string // IP
	Model       string // 功能模块
	ProductNum  string // 产品编号
	Type        string // 操作类型
	Content     string // 内容
	Status      string // 结果
	CreatedAt   string //
	UpdatedAt   string //
	SourceType  string //
	SubUserId   string // 子账号id
	SubUserName string // 子账号用户名
}

//  userLogColumns holds the columns for table yjs_user_log.
var userLogColumns = UserLogColumns{
	Id:          "id",
	UserName:    "user_name",
	Ip:          "ip",
	Model:       "model",
	ProductNum:  "product_num",
	Type:        "type",
	Content:     "content",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	SourceType:  "source_type",
	SubUserId:   "sub_user_id",
	SubUserName: "sub_user_name",
}

// NewUserLogDao creates and returns a new DAO object for table data access.
func NewUserLogDao() *UserLogDao {
	return &UserLogDao{
		group:   "default",
		table:   "yjs_user_log",
		columns: userLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserLogDao) Columns() UserLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
