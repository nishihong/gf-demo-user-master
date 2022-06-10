// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SerialReportCount is the golang structure of table yjs_serial_report_count for DAO operations like Where/Data.
type SerialReportCount struct {
	g.Meta           `orm:"table:yjs_serial_report_count, do:true"`
	Id               interface{} //
	ReportId         interface{} // 序列号
	CountUpLink      interface{} //
	CountDownLink    interface{} //
	CountConnections interface{} //
	Status           interface{} //
	CreatedAt        *gtime.Time //
}
