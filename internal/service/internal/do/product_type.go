// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProductType is the golang structure of table yjs_product_type for DAO operations like Where/Data.
type ProductType struct {
	g.Meta    `orm:"table:yjs_product_type, do:true"`
	Id        interface{} //
	Name      interface{} // 分类名称
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
