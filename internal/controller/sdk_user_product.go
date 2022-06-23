package controller

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/model"

	"github.com/gogf/gf-demo-user/v2/internal/service"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"
)

var SdkUserProduct = cSdkUserProduct{}

type cSdkUserProduct struct{}

func (a *cSdkUserProduct) GetList(ctx context.Context, req *v1.GetSdkListReq) (res *v1.GetSdkListRes, err error) {

	Result, err := service.SdkUserProduct().GetSdkList(ctx, model.GetSdkUserProductInput{
		Keyword:       req.Keyword,
		ProductType:   req.ProductType,
		ProductStatus: req.ProductStatus,
		Search:        req.Search,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		Page:          req.Page,
		Limit:         req.Limit,
	})
	//fmt.Println(Result)

	res = &v1.GetSdkListRes{
		Data:        Result.Data,
		CurrentPage: Result.CurrentPage,
		Total:       Result.Total,
		PerPage:     Result.PerPage,
	}

	return
}

func (a *cSdkUserProduct) Download(ctx context.Context, req *v1.DownloadReq) (res *v1.DownloadRes, err error) {

	res = &v1.DownloadRes{
		DownloadItem: service.SdkUserProduct().Download(ctx),
	}

	return
}
