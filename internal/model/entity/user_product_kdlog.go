// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserProductKdlog is the golang structure for table user_product_kdlog.
type UserProductKdlog struct {
	Id               uint        `json:"id"               ` //
	UserId           int         `json:"userId"           ` //
	UserProductId    int         `json:"userProductId"    ` // 产品id
	Type             int         `json:"type"             ` // 1,是运行扣点，2是封装扣点
	UserSerialNumber string      `json:"userSerialNumber" ` // 用户key值
	OldData          int         `json:"oldData"          ` // 老数据
	NewData          int         `json:"newData"          ` // 新数据
	Mac              string      `json:"mac"              ` //
	Ip               string      `json:"ip"               ` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` //
	CreatedAt        *gtime.Time `json:"createdAt"        ` //
}
