// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SdkProductType is the golang structure for table sdk_product_type.
type SdkProductType struct {
	Id        int         `json:"id"        ` //
	Name      string      `json:"name"      ` // 分类名称
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
