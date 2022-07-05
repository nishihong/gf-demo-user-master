package controller

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/service"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"
)

var SdkProduct = cSdkProduct{}

type cSdkProduct struct{}

func (a *cSdkProduct) GetSdkProductList(ctx context.Context, req *v1.GetSdkProductReq) (res *v1.GetSdkProductRes, err error) {

	Result, err := service.SdkProduct().GetSdkProductList(ctx)

	if err != nil {
		return nil, err
	}

	res = &v1.GetSdkProductRes{
		List: Result,
	}
	return
}
