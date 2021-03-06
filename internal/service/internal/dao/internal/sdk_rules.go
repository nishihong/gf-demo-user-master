// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SdkRulesDao is the data access object for table yjs_sdk_rules.
type SdkRulesDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SdkRulesColumns // columns contains all the column names of Table for convenient usage.
}

// SdkRulesColumns defines and stores column names for table yjs_sdk_rules.
type SdkRulesColumns struct {
	Id            string //
	UserProductId string // 用户产品id
	Type          string // 类型 1 PC端  2 移动端
	CustomIp      string // 客户连接ip
	Port          string // 转发端口
	SourceIp      string // 源站ip
	Status        string // 是否启用
	CreatedAt     string // 创建时间
	UpdatedAt     string //
	DeletedAt     string //
	PortInterval  string // 步长
	PortType      string // 端口类型1单一端口，2多端口（标识0-100样式）
	PortLimit     string // 端口限制个数
	GetIpWay      string // 获取真实ip方式：1不需要，2通过代理，3通过连接
	UserSerialId  string // 用户序列号id
	CustomPort    string // 客户端监听端口
}

//  sdkRulesColumns holds the columns for table yjs_sdk_rules.
var sdkRulesColumns = SdkRulesColumns{
	Id:            "id",
	UserProductId: "user_product_id",
	Type:          "type",
	CustomIp:      "custom_ip",
	Port:          "port",
	SourceIp:      "source_ip",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	PortInterval:  "port_interval",
	PortType:      "port_type",
	PortLimit:     "port_limit",
	GetIpWay:      "get_ip_way",
	UserSerialId:  "user_serial_id",
	CustomPort:    "custom_port",
}

// NewSdkRulesDao creates and returns a new DAO object for table data access.
func NewSdkRulesDao() *SdkRulesDao {
	return &SdkRulesDao{
		group:   "default",
		table:   "yjs_sdk_rules",
		columns: sdkRulesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SdkRulesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SdkRulesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SdkRulesDao) Columns() SdkRulesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SdkRulesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SdkRulesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SdkRulesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
