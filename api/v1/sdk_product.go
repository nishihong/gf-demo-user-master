package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetSdkProductReq struct {
	g.Meta `path:"/api/sdk/product" method:"get" tags:"SdkProduct" summary:"Get the sdk_product list"`
}

//type GetProductRes struct {
//	//*entity.Product
//}

//GetProductRes 查询列表结果
type GetSdkProductRes struct {
	List []SdkProductItem `json:"list"   dc:"列表"`
}

type SdkProductItem struct {
	Id               uint   `json:"id"`
	Name             string `json:"name"`
	Amount           string `json:"amount"`
	SourceServersNum string `json:"source_servers_num"`
	Bandwidth        string `json:"bandwidth"`
	DailyActivate    string `json:"daily_activate"`
}
