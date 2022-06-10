// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-06-06 15:21:14
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PreferentialStrategy1 is the golang structure for table preferential_strategy1.
type PreferentialStrategy1 struct {
	Id                int         `json:"id"                ` //
	Type              int         `json:"type"              ` // 类型 1 购买 2 续费
	Operator          string      `json:"operator"          ` // 操作人
	PreferentialHalf  float64     `json:"preferentialHalf"  ` // 半年折扣
	PreferentialOne   float64     `json:"preferentialOne"   ` // 一年折扣
	PreferentialTwo   float64     `json:"preferentialTwo"   ` // 两年折扣
	PreferentialThree float64     `json:"preferentialThree" ` // 三年折扣
	CreatedAt         *gtime.Time `json:"createdAt"         ` //
	UpdatedAt         *gtime.Time `json:"updatedAt"         ` //
}
