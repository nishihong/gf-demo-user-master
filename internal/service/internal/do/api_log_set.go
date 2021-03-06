// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ApiLogSet is the golang structure of table yjs_api_log_set for DAO operations like Where/Data.
type ApiLogSet struct {
	g.Meta    `orm:"table:yjs_api_log_set, do:true"`
	Id        interface{} //
	OnOff     interface{} // 是否记录结果，1是，2否
	ApiType   interface{} // 需要记录的api协议
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
