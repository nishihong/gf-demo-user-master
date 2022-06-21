package controller

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"
)

var Product = cProduct{}

type cProduct struct{}

func (a *cProduct) GetSdkProductList(ctx context.Context, req *v1.GetStatusReq) (res *v1.GetStatusRes, err error) {

	res = &v1.GetStatusRes{
		List: map[int]string{
			1: "正常",
			2: "已到期",
			3: "禁用",
		},
	}

	return
}
