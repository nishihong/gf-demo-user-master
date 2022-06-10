// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table yjs_user for DAO operations like Where/Data.
type User struct {
	g.Meta           `orm:"table:yjs_user, do:true"`
	Id               interface{} //
	Name             interface{} //
	Email            interface{} //
	Mobile           interface{} //
	Password         interface{} //
	UserId           interface{} //
	RememberToken    interface{} //
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
	IsAdmin          interface{} //
	UserMoney        interface{} // 用户余额
	LoginTime        *gtime.Time //
	LoginIp          interface{} // ip
	LoginInfo        interface{} // 登录信息
	WxKey            interface{} // 微信
	NoShowCodeTime   interface{} // 微信扫码今日是否显示(记录时间)
	NoShowCodeStatus interface{} // 是否显示扫码（永久不显示1）
}
