// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SdkUserProductConfigcommon is the golang structure of table yjs_sdk_user_product_configcommon for DAO operations like Where/Data.
type SdkUserProductConfigcommon struct {
	g.Meta        `orm:"table:yjs_sdk_user_product_configcommon, do:true"`
	Id            interface{} //
	Name          interface{} // 配置名称
	UserProductId interface{} //
	Lip           interface{} // Local IP即终端代理要监听的本地Ip
	Port          interface{} // 端口号
	Guid          interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
