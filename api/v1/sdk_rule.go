package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type GetRuleTypeReq struct {
	g.Meta `path:"/api/sdk/getRuleType" method:"get" tags:"rule" summary:"Get the rule type list"`
}

type GetRuleTypeRes struct {
	Data map[int]string `json:"data"   dc:"列表"`
}

type GetUserIpReq struct {
	g.Meta `path:"/api/sdk/getUserIp" method:"get" tags:"rule" summary:"Get the user ip list"`
}

type GetUserIpRes struct {
	Pc     []string `json:"pc"   dc:"列表"`
	Mobile []string `json:"mobile"   dc:"列表"`
}

type GetSourceServersNumReq struct {
	g.Meta       `path:"/api/sdk/rule/getSourceServersNum" method:"get" tags:"rule" summary:"Get SourceServersNum"`
	UserSerialId int `json:"user_serial_id" v:"required"`
}

type GetSourceServersNumRes struct {
	Data string `json:"data"   dc:"防护信息"`
}

type GetSdkRuleListReq struct {
	g.Meta       `path:"/api/sdk/rule" method:"get" tags:"rule" summary:"Get the rule list"`
	UserSerialId int    `json:"user_serial_id" v:"required"`
	Type         int    `json:"type"`
	Keyword      string `json:"keyword"`
	Port         string `json:"port"`
	Page         int    `json:"page"`
	Limit        int    `json:"limit"`
}

type GetSdkRuleListRes struct {
	Data        []SdkRuleListItem `json:"data"   dc:"列表"`
	CurrentPage int               `json:"current_page"`
	Total       int               `json:"total"`
	PerPage     int               `json:"per_page"`
}

type SdkRuleListItem struct {
	Id            int         `json:"id"`
	UserProductId int         `json:"user_product_id"`
	Type          int         `json:"type"`
	CustomIp      string      `json:"custom_ip"`
	Port          int         `json:"port"`
	SourceIp      string      `json:"source_ip"`
	Status        int         `json:"status"`
	CreatedAt     *gtime.Time `json:"created_at"`
	UpdatedAt     *gtime.Time `json:"updated_at"`
	PortInterval  int         `json:"port_interval"`
	PortType      int         `json:"port_type"`
	PortLimit     int         `json:"port_limit"`
	GetIpWay      int         `json:"get_ip_way"`
	UserSerialId  int         `json:"user_serial_id"`
	TypeName      string      `json:"type_name"`
	PortIntervals string      `json:"port_intervals"`
}

//type EditRuleReq struct {
//	g.Meta `path:"/api/sdk/rule/one" method:"post" tags:"rule" summary:"EditUserSerial"`
//	Id     int    `json:"id" v:"required"`
//	Name   string `json:"name" v:"required"`
//}
//type EditRuleRes struct {
//}
//
//type CreateRuleReq struct {
//	g.Meta           `path:"/api/sdk/rule" method:"post" tags:"rule" summary:"Create"`
//	Name             string `json:"name" v:"required"`
//	SdkUserProductId int    `json:"sdk_user_product_id" v:"required"`
//}

type CreateRuleRes struct {
}

type DeleteRuleReq struct {
	g.Meta       `path:"/api/sdk/rule/delete" method:"post" tags:"rule" summary:"delete"`
	Ids          string `json:"ids" v:"required"`
	UserSerialId int    `json:"user_serial_id" v:"required"`
}

type DeleteRuleRes struct {
}
