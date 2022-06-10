// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SerialReportAvgDao is the data access object for table yjs_serial_report_avg.
type SerialReportAvgDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns SerialReportAvgColumns // columns contains all the column names of Table for convenient usage.
}

// SerialReportAvgColumns defines and stores column names for table yjs_serial_report_avg.
type SerialReportAvgColumns struct {
	Id                 string //
	ProductNum         string //
	ServerIp           string //
	ConnectionsAvg     string //
	ConnectionsMax     string //
	ConnectionsMin     string //
	UpLinkAvg          string //
	UpLinkMax          string //
	UpLinkMin          string //
	DownLinkAvg        string //
	DownLinkMax        string //
	DownLinkMin        string //
	AvgDate            string //
	Type               string // 1-1天 2-7天 3-30天
	CreatedAt          string //
	UpdatedAt          string //
	SdkProductSerialId string // sdk_user_product_id 或 user_serial_id
	ClassType          string // 统计类型 1-产品 2-子序列号
}

//  serialReportAvgColumns holds the columns for table yjs_serial_report_avg.
var serialReportAvgColumns = SerialReportAvgColumns{
	Id:                 "id",
	ProductNum:         "product_num",
	ServerIp:           "server_ip",
	ConnectionsAvg:     "connections_avg",
	ConnectionsMax:     "connections_max",
	ConnectionsMin:     "connections_min",
	UpLinkAvg:          "up_link_avg",
	UpLinkMax:          "up_link_max",
	UpLinkMin:          "up_link_min",
	DownLinkAvg:        "down_link_avg",
	DownLinkMax:        "down_link_max",
	DownLinkMin:        "down_link_min",
	AvgDate:            "avg_date",
	Type:               "type",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	SdkProductSerialId: "sdk_product_serial_id",
	ClassType:          "class_type",
}

// NewSerialReportAvgDao creates and returns a new DAO object for table data access.
func NewSerialReportAvgDao() *SerialReportAvgDao {
	return &SerialReportAvgDao{
		group:   "default",
		table:   "yjs_serial_report_avg",
		columns: serialReportAvgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SerialReportAvgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SerialReportAvgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SerialReportAvgDao) Columns() SerialReportAvgColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SerialReportAvgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SerialReportAvgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SerialReportAvgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
