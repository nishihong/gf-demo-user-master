// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BlackListIp is the golang structure for table black_list_ip.
type BlackListIp struct {
	Id            uint        `json:"id"            ` //
	Ip            string      `json:"ip"            ` //
	CreatedAt     *gtime.Time `json:"createdAt"     ` //
	UpdatedAt     *gtime.Time `json:"updatedAt"     ` //
	UserProductId int         `json:"userProductId" ` //
}
