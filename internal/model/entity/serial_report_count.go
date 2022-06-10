// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SerialReportCount is the golang structure for table serial_report_count.
type SerialReportCount struct {
	Id               uint        `json:"id"               ` //
	ReportId         string      `json:"reportId"         ` // 序列号
	CountUpLink      int         `json:"countUpLink"      ` //
	CountDownLink    int         `json:"countDownLink"    ` //
	CountConnections int         `json:"countConnections" ` //
	Status           int         `json:"status"           ` //
	CreatedAt        *gtime.Time `json:"createdAt"        ` //
}