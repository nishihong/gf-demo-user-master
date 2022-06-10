// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLog is the golang structure of table yjs_user_log for DAO operations like Where/Data.
type UserLog struct {
	g.Meta      `orm:"table:yjs_user_log, do:true"`
	Id          interface{} //
	UserName    interface{} // 用户名
	Ip          interface{} // IP
	Model       interface{} // 功能模块
	ProductNum  interface{} // 产品编号
	Type        interface{} // 操作类型
	Content     interface{} // 内容
	Status      interface{} // 结果
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	SourceType  interface{} //
	SubUserId   interface{} // 子账号id
	SubUserName interface{} // 子账号用户名
}