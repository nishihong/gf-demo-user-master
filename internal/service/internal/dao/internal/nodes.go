// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NodesDao is the data access object for table yjs_nodes.
type NodesDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns NodesColumns // columns contains all the column names of Table for convenient usage.
}

// NodesColumns defines and stores column names for table yjs_nodes.
type NodesColumns struct {
	Id               string //
	Name             string // 节点名称
	NodeType         string // 1防御网关。2 安全网关
	Type             string // 类型 1是共享 2是独立
	Status           string // 1是启用 2是不启用
	Port             string // 端口号
	CreatedAt        string //
	UpdatedAt        string //
	State            string // 是否启用 1正常2禁用3离线
	UserSerialId     string // 分配的序列号id，可为空
	DistributionType string // 1手动分配，2自动分配
	NodeVersion      string // 节点版本
	DistType         string // 分配类型 1、实例ID，2、序列号
	Timer            string // 运行时间
	NsThreshold      string // 次数
	UploadTimer      string // 上报时间
	Area             string // 地区配置
	AreaSafe         string // 地区配置-用于匹配安全网关
	AreaLine         string // 专线ip-地址
	IpLine           string // 专线IP
}

//  nodesColumns holds the columns for table yjs_nodes.
var nodesColumns = NodesColumns{
	Id:               "id",
	Name:             "name",
	NodeType:         "node_type",
	Type:             "type",
	Status:           "status",
	Port:             "port",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	State:            "state",
	UserSerialId:     "user_serial_id",
	DistributionType: "distribution_type",
	NodeVersion:      "node_version",
	DistType:         "dist_type",
	Timer:            "timer",
	NsThreshold:      "ns_threshold",
	UploadTimer:      "upload_timer",
	Area:             "area",
	AreaSafe:         "area_safe",
	AreaLine:         "area_line",
	IpLine:           "ip_line",
}

// NewNodesDao creates and returns a new DAO object for table data access.
func NewNodesDao() *NodesDao {
	return &NodesDao{
		group:   "default",
		table:   "yjs_nodes",
		columns: nodesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NodesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NodesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NodesDao) Columns() NodesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NodesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NodesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NodesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
