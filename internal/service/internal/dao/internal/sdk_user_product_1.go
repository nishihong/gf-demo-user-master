// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SdkUserProduct1Dao is the data access object for table yjs_sdk_user_product1.
type SdkUserProduct1Dao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns SdkUserProduct1Columns // columns contains all the column names of Table for convenient usage.
}

// SdkUserProduct1Columns defines and stores column names for table yjs_sdk_user_product1.
type SdkUserProduct1Columns struct {
	Id                  string //
	UserId              string // 用户id
	ProductId           string // 产品id
	UserSerialId        string // 用户序列号表id
	ProductType         string // 对应产品表的类型
	HostName            string // 1是服务器2是云
	HostType            string // 1是服务器2是云
	HostId              string // 产品id
	OnOff               string // -1是未初始化，1是启用，2是关闭
	Status              string // 用户产品状态1正常 2关闭 -1未支付
	Name                string // 产品名称
	AvailablePoints     string // 可用点数
	AvaliableFrequency  string // 可用封装次数
	CreatedAt           string //
	UpdatedAt           string //
	StartAt             string // 生效日期
	EndAt               string // 失效日期
	Bz                  string //
	ClientLimit         string // 客户端最大连接数
	SharedNodesNum      string // 共享节点数
	IndependentNodesNum string // 独立节点数
	CountUpLink         string // 统计数据
	CountDownLink       string // 统计数据
	CountConnections    string // 统计数据
	CountAt             string // 统计数据时间
	IsLogo              string // 是否展示logo,1展示，0不展示
	Username            string // sdk字段-所属用户
	ProductNum          string // sdk字段-产品编号
	Amount              string // sdk字段-金额
	TradedMode          string // sdk字段-交易方式
	TradedType          string // sdk字段-交易类型
	OrderType           string // sdk字段-订单类型
	DeletedAt           string //
	Operator            string // SDK字段
	SourceType          string // 客户来源 1云加速 2sdk
	SdkNum              string // sdk_num
	Guid                string // 规则改变标识
	SkinFile            string // 客户端皮肤资源包 文件ID
	IcoImg              string // 托盘图标 文件ID
	Remarks             string // 数据中心-列表备注
	RealAmount          string // 数据中心-实际支付价格
	SellId              string // 销售id
	NodeLimit           string // 节点下发数量限制
	GameNum             string // 游戏数量限制
}

//  sdkUserProduct1Columns holds the columns for table yjs_sdk_user_product1.
var sdkUserProduct1Columns = SdkUserProduct1Columns{
	Id:                  "id",
	UserId:              "user_id",
	ProductId:           "product_id",
	UserSerialId:        "user_serial_id",
	ProductType:         "product_type",
	HostName:            "host_name",
	HostType:            "host_type",
	HostId:              "host_id",
	OnOff:               "on_off",
	Status:              "status",
	Name:                "name",
	AvailablePoints:     "available_points",
	AvaliableFrequency:  "avaliable_frequency",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	StartAt:             "start_at",
	EndAt:               "end_at",
	Bz:                  "bz",
	ClientLimit:         "client_limit",
	SharedNodesNum:      "shared_nodes_num",
	IndependentNodesNum: "independent_nodes_num",
	CountUpLink:         "count_up_link",
	CountDownLink:       "count_down_link",
	CountConnections:    "count_connections",
	CountAt:             "count_at",
	IsLogo:              "is_logo",
	Username:            "username",
	ProductNum:          "product_num",
	Amount:              "amount",
	TradedMode:          "traded_mode",
	TradedType:          "traded_type",
	OrderType:           "order_type",
	DeletedAt:           "deleted_at",
	Operator:            "operator",
	SourceType:          "source_type",
	SdkNum:              "sdk_num",
	Guid:                "guid",
	SkinFile:            "skin_file",
	IcoImg:              "ico_img",
	Remarks:             "remarks",
	RealAmount:          "real_amount",
	SellId:              "sell_id",
	NodeLimit:           "node_limit",
	GameNum:             "game_num",
}

// NewSdkUserProduct1Dao creates and returns a new DAO object for table data access.
func NewSdkUserProduct1Dao() *SdkUserProduct1Dao {
	return &SdkUserProduct1Dao{
		group:   "default",
		table:   "yjs_sdk_user_product1",
		columns: sdkUserProduct1Columns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SdkUserProduct1Dao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SdkUserProduct1Dao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SdkUserProduct1Dao) Columns() SdkUserProduct1Columns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SdkUserProduct1Dao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SdkUserProduct1Dao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SdkUserProduct1Dao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}