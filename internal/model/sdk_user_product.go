package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type GetSdkUserProductInput struct {
	Keyword       string      `json:"keyword"`
	ProductType   int         `json:"product_type"`
	ProductStatus int         `json:"product_status"`
	Search        string      `json:"search"`
	StartTime     *gtime.Time `json:"start_time"`
	EndTime       *gtime.Time `json:"end_time"`
	Page          int         `json:"page" d:"1"`
	Limit         int         `json:"limit" d:"10"`
}

// 组合模型
//type Entity struct {
//	SdkUserProduct        *EntitySdkUserProduct
//	SdkUserProductSetting *EntitySdkUserProductSetting
//	SdkUserSerial         *SdkUserSerial
//}
