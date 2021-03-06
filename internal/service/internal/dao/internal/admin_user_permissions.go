// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminUserPermissionsDao is the data access object for table yjs_admin_user_permissions.
type AdminUserPermissionsDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns AdminUserPermissionsColumns // columns contains all the column names of Table for convenient usage.
}

// AdminUserPermissionsColumns defines and stores column names for table yjs_admin_user_permissions.
type AdminUserPermissionsColumns struct {
	UserId       string //
	PermissionId string //
	CreatedAt    string //
	UpdatedAt    string //
}

//  adminUserPermissionsColumns holds the columns for table yjs_admin_user_permissions.
var adminUserPermissionsColumns = AdminUserPermissionsColumns{
	UserId:       "user_id",
	PermissionId: "permission_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewAdminUserPermissionsDao creates and returns a new DAO object for table data access.
func NewAdminUserPermissionsDao() *AdminUserPermissionsDao {
	return &AdminUserPermissionsDao{
		group:   "default",
		table:   "yjs_admin_user_permissions",
		columns: adminUserPermissionsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminUserPermissionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminUserPermissionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminUserPermissionsDao) Columns() AdminUserPermissionsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminUserPermissionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminUserPermissionsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminUserPermissionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
