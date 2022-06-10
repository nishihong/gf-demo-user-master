// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SdkBlackListMacDao is the data access object for table yjs_sdk_black_list_mac.
type SdkBlackListMacDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns SdkBlackListMacColumns // columns contains all the column names of Table for convenient usage.
}

// SdkBlackListMacColumns defines and stores column names for table yjs_sdk_black_list_mac.
type SdkBlackListMacColumns struct {
	Id            string //
	Mac           string //
	CreatedAt     string //
	UpdatedAt     string //
	UserProductId string //
}

//  sdkBlackListMacColumns holds the columns for table yjs_sdk_black_list_mac.
var sdkBlackListMacColumns = SdkBlackListMacColumns{
	Id:            "id",
	Mac:           "mac",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	UserProductId: "user_product_id",
}

// NewSdkBlackListMacDao creates and returns a new DAO object for table data access.
func NewSdkBlackListMacDao() *SdkBlackListMacDao {
	return &SdkBlackListMacDao{
		group:   "default",
		table:   "yjs_sdk_black_list_mac",
		columns: sdkBlackListMacColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SdkBlackListMacDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SdkBlackListMacDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SdkBlackListMacDao) Columns() SdkBlackListMacColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SdkBlackListMacDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SdkBlackListMacDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SdkBlackListMacDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}