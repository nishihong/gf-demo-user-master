package v1

import (
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetProductReq struct {
	g.Meta `path:"/api/sdk/product" method:"get" tags:"Product" summary:"Get the product list"`
}
type GetProductRes struct {
	*entity.Product
}
