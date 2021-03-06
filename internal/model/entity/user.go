// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id               uint        `json:"id"               ` //
	Name             string      `json:"name"             ` //
	Email            string      `json:"email"            ` //
	Mobile           string      `json:"mobile"           ` //
	Password         string      `json:"password"         ` //
	UserId           int         `json:"userId"           ` //
	RememberToken    string      `json:"rememberToken"    ` //
	CreatedAt        *gtime.Time `json:"createdAt"        ` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` //
	IsAdmin          int         `json:"isAdmin"          ` //
	UserMoney        float64     `json:"userMoney"        ` // 用户余额
	LoginTime        *gtime.Time `json:"loginTime"        ` //
	LoginIp          string      `json:"loginIp"          ` // ip
	LoginInfo        string      `json:"loginInfo"        ` // 登录信息
	WxKey            string      `json:"wxKey"            ` // 微信
	NoShowCodeTime   string      `json:"noShowCodeTime"   ` // 微信扫码今日是否显示(记录时间)
	NoShowCodeStatus int         `json:"noShowCodeStatus" ` // 是否显示扫码（永久不显示1）
}
