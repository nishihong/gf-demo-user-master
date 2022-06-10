// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SerialReportCountIp is the golang structure for table serial_report_count_ip.
type SerialReportCountIp struct {
	Id               uint        `json:"id"               ` //
	ServerIp         string      `json:"serverIp"         ` // ip
	CountUpLink      int         `json:"countUpLink"      ` //
	CountDownLink    int         `json:"countDownLink"    ` //
	CountConnections int         `json:"countConnections" ` //
	Status           int         `json:"status"           ` //
	CreatedAt        *gtime.Time `json:"createdAt"        ` //
}
