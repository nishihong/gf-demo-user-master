// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LoginLog is the golang structure of table yjs_login_log for DAO operations like Where/Data.
type LoginLog struct {
	g.Meta    `orm:"table:yjs_login_log, do:true"`
	Id        interface{} //
	Username  interface{} //
	UserId    interface{} //
	LoginType interface{} // 登录方式，1云加速，2官网
	LoginIp   interface{} // 登录ip
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
