// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserProductCount is the golang structure of table yjs_user_product_count for DAO operations like Where/Data.
type UserProductCount struct {
	g.Meta           `orm:"table:yjs_user_product_count, do:true"`
	Id               interface{} //
	UserProductCount interface{} //
	YxdCount         interface{} //
	YjsCount         interface{} //
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
}