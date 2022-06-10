// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminUsersDao is the data access object for table yjs_admin_users.
type AdminUsersDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns AdminUsersColumns // columns contains all the column names of Table for convenient usage.
}

// AdminUsersColumns defines and stores column names for table yjs_admin_users.
type AdminUsersColumns struct {
	Id            string //
	Username      string //
	Password      string //
	Name          string //
	Avatar        string //
	RememberToken string //
	CreatedAt     string //
	UpdatedAt     string //
}

//  adminUsersColumns holds the columns for table yjs_admin_users.
var adminUsersColumns = AdminUsersColumns{
	Id:            "id",
	Username:      "username",
	Password:      "password",
	Name:          "name",
	Avatar:        "avatar",
	RememberToken: "remember_token",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewAdminUsersDao creates and returns a new DAO object for table data access.
func NewAdminUsersDao() *AdminUsersDao {
	return &AdminUsersDao{
		group:   "default",
		table:   "yjs_admin_users",
		columns: adminUsersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminUsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminUsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminUsersDao) Columns() AdminUsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminUsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminUsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminUsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
