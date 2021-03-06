// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

// Province is the golang structure for table province.
type Province struct {
	Id           int    `json:"id"           ` // 主键自增
	Pid          int    `json:"pid"          ` // 父类id
	DistrictName string `json:"districtName" ` // 城市的名字
	Type         int    `json:"type"         ` // 城市的类型，0是国，1是省，2是市，3是区
	Hierarchy    int    `json:"hierarchy"    ` // 地区所处的层级
	DistrictSqe  string `json:"districtSqe"  ` // 层级序列
	Partition    int    `json:"partition"    ` // 分区，省有效 0其他 1东北 2华东 3华北 4华中 5华南 6西南 7西北
}
