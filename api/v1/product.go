package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetStatusReq struct {
	g.Meta `path:"/api/sdk/getStatus" method:"get" tags:"." summary:"Get the . list"`
}

type GetStatusRes struct {
	List map[int]string
}
