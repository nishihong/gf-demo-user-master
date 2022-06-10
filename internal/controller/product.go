package controller

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/ghttp"
)

var Product = cProduct{}

type cProduct struct{}

func (a *cProduct) GetProductList(r *ghttp.Request) {
	m := g.DB().Model("product")
	var productList []*cProduct
	m.Where("id>", "1").Scan(&productList)
	r.Response.WriteJsonExit(&productList)
}

//func (c *cProduct) GetProductList(ctx context.Context, req *v1.GetProductReq) (res *v1.GetProductRes, err error) {
//	res = &v1.GetProductRes{
//		Product: service.Product().GetProductList(ctx, "基础版"),
//	}
//	return
//}
