// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// NodeProvince is the golang structure of table yjs_node_province for DAO operations like Where/Data.
type NodeProvince struct {
	g.Meta       `orm:"table:yjs_node_province, do:true"`
	Id           interface{} //
	NodeId       interface{} // 节点id
	Province     interface{} // 省份
	ProvinceType interface{} // 1、省级配置，2、市级配置
	MachineId    interface{} // 机房id
	City         interface{} // 城市
	Line         interface{} // 线路1电信2联通3移动
	CreatedAt    interface{} //
	UpdatedAt    interface{} //
	SerialList   interface{} // 序列号
	Area         interface{} // 省份
	SetType      interface{} // 1允许 2限制
}
