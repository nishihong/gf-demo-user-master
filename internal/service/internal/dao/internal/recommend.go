// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RecommendDao is the data access object for table yjs_recommend.
type RecommendDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns RecommendColumns // columns contains all the column names of Table for convenient usage.
}

// RecommendColumns defines and stores column names for table yjs_recommend.
type RecommendColumns struct {
	Id            string //
	AdminId       string // 发布人ID
	AdminSurename string // 发布人姓名
	Title         string // 图片提示
	Img           string // 图片路径
	Url           string // 图片链接路径
	Sort          string // 排序
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	OnOff         string // 是否开启 1打开2关闭
}

//  recommendColumns holds the columns for table yjs_recommend.
var recommendColumns = RecommendColumns{
	Id:            "id",
	AdminId:       "admin_id",
	AdminSurename: "admin_surename",
	Title:         "title",
	Img:           "img",
	Url:           "url",
	Sort:          "sort",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	OnOff:         "on_off",
}

// NewRecommendDao creates and returns a new DAO object for table data access.
func NewRecommendDao() *RecommendDao {
	return &RecommendDao{
		group:   "default",
		table:   "yjs_recommend",
		columns: recommendColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RecommendDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RecommendDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RecommendDao) Columns() RecommendColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RecommendDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RecommendDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RecommendDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
