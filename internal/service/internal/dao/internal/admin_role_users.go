// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminRoleUsersDao is the data access object for table yjs_admin_role_users.
type AdminRoleUsersDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns AdminRoleUsersColumns // columns contains all the column names of Table for convenient usage.
}

// AdminRoleUsersColumns defines and stores column names for table yjs_admin_role_users.
type AdminRoleUsersColumns struct {
	RoleId    string //
	UserId    string //
	CreatedAt string //
	UpdatedAt string //
}

//  adminRoleUsersColumns holds the columns for table yjs_admin_role_users.
var adminRoleUsersColumns = AdminRoleUsersColumns{
	RoleId:    "role_id",
	UserId:    "user_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAdminRoleUsersDao creates and returns a new DAO object for table data access.
func NewAdminRoleUsersDao() *AdminRoleUsersDao {
	return &AdminRoleUsersDao{
		group:   "default",
		table:   "yjs_admin_role_users",
		columns: adminRoleUsersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminRoleUsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminRoleUsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminRoleUsersDao) Columns() AdminRoleUsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminRoleUsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminRoleUsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminRoleUsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
