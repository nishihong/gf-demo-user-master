package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetUserSerialListReq struct {
	g.Meta `path:"/api/sdk/userSerial" method:"get" tags:"userSerial" summary:"Get the sdk_user_serial list"`
	Id     int `json:"id" v:"required"`
	//Page   int `json:"page"`
	//Limit  int `json:"limit"`
}

type GetUserSerialListRes struct {
	Data []SdkSerialListItem `json:"data"   dc:"列表"`
}

type SdkSerialListItem struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Number    string `json:"number"`
	Name      string `json:"name"`
	HostCount int    `json:"host_count"`
}

type GetKeyReq struct {
	g.Meta `path:"/api/sdk/userSerial/getKey" method:"get" tags:"getKey" summary:"Get key"`
	Id     int `json:"id" v:"required"`
}
type GetKeyRes struct {
	Data string `json:"data"`
}

type EditUserSerialReq struct {
	g.Meta `path:"/api/sdk/userSerial/one" method:"post" tags:"EditUserSerial" summary:"EditUserSerial"`
	Id     int    `json:"id" v:"required"`
	Name   string `json:"name" v:"required"`
}
type EditUserSerialRes struct {
}

type CreateUserSerialReq struct {
	g.Meta           `path:"/api/sdk/userSerial" method:"post" tags:"userSerial" summary:"Create"`
	Name             string `json:"name"`
	SdkUserProductId int    `json:"sdk_user_product_id"`
}

type CreateUserSerialRes struct {
}

type DeleteUserSerialReq struct {
	g.Meta `path:"/api/sdk/userSerial/delete" method:"post" tags:"userSerial" summary:"delete"`
	Id     int `json:"id"`
}

type DeleteUserSerialRes struct {
}
