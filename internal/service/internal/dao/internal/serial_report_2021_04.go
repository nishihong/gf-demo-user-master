// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SerialReport202104Dao is the data access object for table yjs_serial_report_2021-04.
type SerialReport202104Dao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns SerialReport202104Columns // columns contains all the column names of Table for convenient usage.
}

// SerialReport202104Columns defines and stores column names for table yjs_serial_report_2021-04.
type SerialReport202104Columns struct {
	Id          string //
	Port        string //
	ProductNum  string //
	Connections string //
	UpLink      string //
	DownLink    string //
	ServerIp    string //
	CreatedAt   string //
	UpdatedAt   string //
}

//  serialReport202104Columns holds the columns for table yjs_serial_report_2021-04.
var serialReport202104Columns = SerialReport202104Columns{
	Id:          "id",
	Port:        "port",
	ProductNum:  "product_num",
	Connections: "connections",
	UpLink:      "up_link",
	DownLink:    "down_link",
	ServerIp:    "server_ip",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewSerialReport202104Dao creates and returns a new DAO object for table data access.
func NewSerialReport202104Dao() *SerialReport202104Dao {
	return &SerialReport202104Dao{
		group:   "default",
		table:   "yjs_serial_report_2021-04",
		columns: serialReport202104Columns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SerialReport202104Dao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SerialReport202104Dao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SerialReport202104Dao) Columns() SerialReport202104Columns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SerialReport202104Dao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SerialReport202104Dao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SerialReport202104Dao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
