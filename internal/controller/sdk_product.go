package controller

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/service"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"
)

var SdkProduct = cSdkProduct{}

type cSdkProduct struct{}

func (a *cSdkProduct) GetSdkProductList(ctx context.Context, req *v1.GetSdkProductReq) (res *v1.GetSdkProductRes, err error) {

	res = &v1.GetSdkProductRes{
		List: service.SdkProduct().GetSdkProductList(ctx),
	}
	return
}
