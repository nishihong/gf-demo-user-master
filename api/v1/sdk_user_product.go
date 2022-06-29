package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetSdkListReq struct {
	g.Meta        `path:"/api/sdk" method:"get" tags:"sdk_user_product" summary:"Get the sdk_user_product list"`
	Keyword       string      `json:"keyword"`
	ProductType   int         `json:"product_type"`
	ProductStatus int         `json:"product_status"`
	Search        string      `json:"search"`
	StartTime     *gtime.Time `json:"start_time"
`
	EndTime *gtime.Time `json:"end_time"`
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
}

type GetSdkListRes struct {
	Data        []SdkListItem `json:"data"   dc:"列表"`
	CurrentPage int           `json:"current_page"`
	Total       int           `json:"total"`
	PerPage     int           `json:"per_page"`
}

type SdkListItem struct {
	Id              uint        `json:"id"`
	DeviceId        string      `json:"device_id"`
	ProductName     string      `json:"product_name"`
	ProductType     uint        `json:"product_type"`
	ProductTypeName string      `json:"product_type_name"`
	Username        string      `json:"username"`
	Uuid            uint        `json:"uuid"`
	StartTime       *gtime.Time `json:"start_time"`
	EndTime         *gtime.Time `json:"end_time"`
	ProductStatus   string      `json:"product_status"`
}

type DownloadReq struct {
	g.Meta `path:"/api/sdk/download" method:"get" tags:"sdk_user_product" summary:"Get the profile of download"`
}
type DownloadRes struct {
	*DownloadItem
}

type DownloadItem struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	FileName string `json:"file_name"`
	FileUrl  string `json:"file_url"`
}
