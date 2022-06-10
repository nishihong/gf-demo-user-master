// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

// SoftProvince is the golang structure for table soft_province.
type SoftProvince struct {
	Id           uint   `json:"id"           ` //
	SoftId       int    `json:"softId"       ` // 主程序id
	Province     string `json:"province"     ` // 省份
	ProvinceType int    `json:"provinceType" ` // 1、省级配置，2、市级配置
	City         string `json:"city"         ` // 城市
	CreatedAt    string `json:"createdAt"    ` //
	UpdatedAt    string `json:"updatedAt"    ` //
}
