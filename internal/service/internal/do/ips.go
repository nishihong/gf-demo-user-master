// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Ips is the golang structure of table yjs_ips for DAO operations like Where/Data.
type Ips struct {
	g.Meta   `orm:"table:yjs_ips, do:true"`
	Start    interface{} //
	End      interface{} //
	Startip  interface{} //
	Endip    interface{} //
	Country  interface{} //
	Province interface{} //
	City     interface{} //
	Operator interface{} //
	Zipcode  interface{} //
	Areacode interface{} //
}
