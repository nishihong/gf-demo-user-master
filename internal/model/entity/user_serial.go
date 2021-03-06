// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserSerial is the golang structure for table user_serial.
type UserSerial struct {
	Id               uint        `json:"id"               ` // 客户类型
	Lip              string      `json:"lip"              ` // 每一个序列号对应一个lip
	Port             int         `json:"port"             ` // 端口号
	UserId           int         `json:"userId"           ` // 用户id
	Number           string      `json:"number"           ` // 用户序列号
	Remark           string      `json:"remark"           ` // 序列号名称
	ProductType      int         `json:"productType"      ` // 序列号类型 1免费 2普通 3定制
	OnOff            int         `json:"onOff"            ` // 1是启用2是关闭
	CreatedAt        *gtime.Time `json:"createdAt"        ` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` //
	Status           int         `json:"status"           ` // 状态 1正常 2关闭 -1 未使用
	LinkType         string      `json:"linkType"         ` // 1是源站，2是节点
	CustomStatus     int         `json:"customStatus"     ` // 自定义转发ip,1正常，2关闭
	CountUpLink      int         `json:"countUpLink"      ` // 统计数据
	CountDownLink    int         `json:"countDownLink"    ` // 统计数据
	CountConnections int         `json:"countConnections" ` // 统计数据
	CountAt          *gtime.Time `json:"countAt"          ` // 统计数据时间
	DeletedAt        *gtime.Time `json:"deletedAt"        ` //
	SourceType       int         `json:"sourceType"       ` // 客户类型 1云加速 2游戏盾
}
