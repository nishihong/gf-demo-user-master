// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HostVersionDao is the data access object for table yjs_host_version.
type HostVersionDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns HostVersionColumns // columns contains all the column names of Table for convenient usage.
}

// HostVersionColumns defines and stores column names for table yjs_host_version.
type HostVersionColumns struct {
	Id        string //
	HostId    string //
	HostIdMd5 string //
	Version   string //
	CreatedAt string //
	UpdatedAt string //
}

//  hostVersionColumns holds the columns for table yjs_host_version.
var hostVersionColumns = HostVersionColumns{
	Id:        "id",
	HostId:    "host_id",
	HostIdMd5: "host_id_md5",
	Version:   "version",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewHostVersionDao creates and returns a new DAO object for table data access.
func NewHostVersionDao() *HostVersionDao {
	return &HostVersionDao{
		group:   "default",
		table:   "yjs_host_version",
		columns: hostVersionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *HostVersionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *HostVersionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *HostVersionDao) Columns() HostVersionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *HostVersionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *HostVersionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *HostVersionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
