// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NodeProvinceDao is the data access object for table yjs_node_province.
type NodeProvinceDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns NodeProvinceColumns // columns contains all the column names of Table for convenient usage.
}

// NodeProvinceColumns defines and stores column names for table yjs_node_province.
type NodeProvinceColumns struct {
	Id           string //
	NodeId       string // 节点id
	Province     string // 省份
	ProvinceType string // 1、省级配置，2、市级配置
	MachineId    string // 机房id
	City         string // 城市
	Line         string // 线路1电信2联通3移动
	CreatedAt    string //
	UpdatedAt    string //
	SerialList   string // 序列号
	Area         string // 省份
	SetType      string // 1允许 2限制
}

//  nodeProvinceColumns holds the columns for table yjs_node_province.
var nodeProvinceColumns = NodeProvinceColumns{
	Id:           "id",
	NodeId:       "node_id",
	Province:     "province",
	ProvinceType: "province_type",
	MachineId:    "machine_id",
	City:         "city",
	Line:         "line",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	SerialList:   "serial_list",
	Area:         "area",
	SetType:      "set_type",
}

// NewNodeProvinceDao creates and returns a new DAO object for table data access.
func NewNodeProvinceDao() *NodeProvinceDao {
	return &NodeProvinceDao{
		group:   "default",
		table:   "yjs_node_province",
		columns: nodeProvinceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NodeProvinceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NodeProvinceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NodeProvinceDao) Columns() NodeProvinceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NodeProvinceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NodeProvinceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NodeProvinceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}