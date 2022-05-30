package controller

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"
)

var Product = cProduct{}

type cProduct struct{}

func (c *cProduct) getProduct(ctx context.Context, req *v1.GetProductReq) (res *v1.GetProductRes, err error) {
	res = &v1.GetProductRes{
		cProduct: service.cProduct().GetProfile(ctx),
	}
	return
}
