// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SerialReport202104 is the golang structure of table yjs_serial_report_2021-04 for DAO operations like Where/Data.
type SerialReport202104 struct {
	g.Meta      `orm:"table:yjs_serial_report_2021-04, do:true"`
	Id          interface{} //
	Port        interface{} //
	ProductNum  interface{} //
	Connections interface{} //
	UpLink      interface{} //
	DownLink    interface{} //
	ServerIp    interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
