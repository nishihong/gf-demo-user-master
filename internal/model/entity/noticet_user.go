// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticetUser is the golang structure for table noticet_user.
type NoticetUser struct {
	Id        int         `json:"id"        ` //
	UserId    int         `json:"userId"    ` // 用户id
	NoticetId int         `json:"noticetId" ` // 公告id
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
