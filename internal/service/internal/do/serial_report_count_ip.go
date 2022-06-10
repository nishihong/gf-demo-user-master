// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SerialReportCountIp is the golang structure of table yjs_serial_report_count_ip for DAO operations like Where/Data.
type SerialReportCountIp struct {
	g.Meta           `orm:"table:yjs_serial_report_count_ip, do:true"`
	Id               interface{} //
	ServerIp         interface{} // ip
	CountUpLink      interface{} //
	CountDownLink    interface{} //
	CountConnections interface{} //
	Status           interface{} //
	CreatedAt        *gtime.Time //
}
