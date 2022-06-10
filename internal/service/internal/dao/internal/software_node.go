// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SoftwareNodeDao is the data access object for table yjs_software_node.
type SoftwareNodeDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SoftwareNodeColumns // columns contains all the column names of Table for convenient usage.
}

// SoftwareNodeColumns defines and stores column names for table yjs_software_node.
type SoftwareNodeColumns struct {
	Id        string //
	FileId    string // 对应上传文件id
	FilePath  string //
	Name      string // 软件名称
	Version   string // 版本号
	Remark    string // 备注
	Sort      string // 排序
	CreatedAt string //
	UpdatedAt string //
	Type      string // 1是客户端，2是服务的，3是合集
}

//  softwareNodeColumns holds the columns for table yjs_software_node.
var softwareNodeColumns = SoftwareNodeColumns{
	Id:        "id",
	FileId:    "file_id",
	FilePath:  "file_path",
	Name:      "name",
	Version:   "version",
	Remark:    "remark",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Type:      "type",
}

// NewSoftwareNodeDao creates and returns a new DAO object for table data access.
func NewSoftwareNodeDao() *SoftwareNodeDao {
	return &SoftwareNodeDao{
		group:   "default",
		table:   "yjs_software_node",
		columns: softwareNodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SoftwareNodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SoftwareNodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SoftwareNodeDao) Columns() SoftwareNodeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SoftwareNodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SoftwareNodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SoftwareNodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
