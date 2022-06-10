// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SoftwarePmain is the golang structure for table software_pmain.
type SoftwarePmain struct {
	Id                     int         `json:"id"                     ` //
	FileId                 int         `json:"fileId"                 ` // 对应上传文件id
	FilePath               string      `json:"filePath"               ` //
	ProductType            int         `json:"productType"            ` // 序列号类型 1免费 2普通 3定制
	Name                   string      `json:"name"                   ` // 软件名称
	Version                string      `json:"version"                ` // 版本号
	Remark                 string      `json:"remark"                 ` // 备注
	Sort                   int         `json:"sort"                   ` // 排序
	OperatingSystemVersion string      `json:"operatingSystemVersion" ` // 操作系统版本
	OperatingSystemNum     string      `json:"operatingSystemNum"     ` // 操作系统位数
	CreatedAt              *gtime.Time `json:"createdAt"              ` //
	UpdatedAt              *gtime.Time `json:"updatedAt"              ` //
}
