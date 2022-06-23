package service

import (
	"context"
	"fmt"

	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao"

	"github.com/gogf/gf-demo-user/v2/internal/model"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	// sSdkProduct is service struct of module Product.
	sSdkUserProduct struct{}
)

var (
	// insSdkProduct is the instance of service Product.
	insSdkUserProduct = sSdkUserProduct{}
)

// Product returns the interface of Product service.
func SdkUserProduct() *sSdkUserProduct {
	return &insSdkUserProduct
}

// GetSdkList
func (s *sSdkUserProduct) GetSdkList(ctx context.Context, in model.GetSdkUserProductInput) (list v1.GetSdkListRes, err error) {
	var (
		likePattern = `%` + in.Keyword + `%`
	)

	//默认值
	if in.Page == 0 {
		in.Page = 1
	}
	if in.Limit == 0 {
		in.Limit = 10
	}

	list = v1.GetSdkListRes{
		CurrentPage: in.Page,
		PerPage:     in.Limit,
	}

	//默认方法
	mod := g.Model("yjs_sdk_user_product", "a").
		LeftJoin("yjs_sdk_user_product_setting b", "a.id=b.user_product_id")

	//条件
	//if in.UserId != "" {
	//	m = m.Where("a.user_id", in.UserId)
	//}
	if in.Keyword != "" {
		mod = mod.WhereLike("a.device_id", likePattern).WhereOrLike("a.bz", likePattern)
	}
	if in.ProductType != 0 {
		mod = mod.Where("a.product_id=", in.ProductType)
	}
	if in.ProductStatus != 0 {
		if in.ProductStatus == 1 {
			mod = mod.Where("a.on_off=", in.ProductStatus)
			mod = mod.Where("a.end_at>", gtime.Datetime())
		} else {
			mod = mod.Where("a.on_off=", in.ProductStatus)
			mod = mod.Where("a.end_at<", gtime.Datetime())
		}
	}
	if in.StartTime != nil && in.EndTime != nil {
		var (
			startTime = gtime.New((in.StartTime)).Format("Y-m-d H:i:s")
			EndTime   = gtime.New((in.StartTime)).EndOfDay().Format("Y-m-d H:i:s")
		)
		mod = mod.WhereBetween("a.end_at", startTime, EndTime)
	} else {
		if in.StartTime != nil {
			mod = mod.Where("a.start_at>=", gtime.New((in.StartTime)).Format("Y-m-d H:i:s"))
		} else if in.EndTime != nil {
			mod = mod.Where("a.end_at<=", gtime.New((in.EndTime)).EndOfDay().Format("Y-m-d H:i:s"))
		}
	}

	list.Total, err = mod.Count()
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return list, err
	}

	res, err := mod.Fields("a.id,a.device_id,a.bz,a.product_id,a.username,a.user_id,a.start_at,a.end_at,a.on_off,a.bz,b.name").
		Order("a.id asc").
		Page(in.Page, in.Limit).
		All()
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return list, err
	}

	//定义长度 不先定义会有越界的问题
	results := make([]v1.SdkListItem, len(res))

	//数据处理
	for i, v := range res {
		//fmt.Println(i, v)
		product_status := "关闭"
		if v["on_off"].Int() == 1 {
			product_status = "正常"
		}

		results[i].Id = v["id"].Uint()
		results[i].DeviceId = v["device_id"].String()
		results[i].ProductName = v["bz"].String()
		results[i].ProductType = v["product_id"].Uint()
		results[i].ProductTypeName = v["name"].String()
		results[i].Username = v["username"].String()
		results[i].Uuid = v["user_id"].Uint()
		results[i].StartTime = v["start_at"].GTime()
		results[i].EndTime = v["end_at"].GTime()
		results[i].ProductStatus = product_status
	}
	list.Data = results
	//fmt.Println(list)

	return list, nil
}

// Download.
func (s *sSdkUserProduct) Download(ctx context.Context) (res *v1.DownloadItem) {

	mod := dao.SdkSoftware.Ctx(ctx)

	sdkSoftwareInfo, err := mod.Where("type=", 5).Order("id desc").One()
	fmt.Println(sdkSoftwareInfo)
	if err != nil {
		err = gerror.Newf(`ErrorORM`)

		return nil
	}

	fileInfo, err := dao.File.Ctx(ctx).Where("id=", sdkSoftwareInfo["file_id"]).One()
	if err != nil {
		err = gerror.Newf(`ErrorORM`)

		return nil
	}

	//处理数据
	res = &v1.DownloadItem{
		Id:       sdkSoftwareInfo["id"].Int(),
		Name:     sdkSoftwareInfo["name"].String(),
		FileName: fileInfo["file_name"].String(),
		FileUrl:  fileInfo["file_url"].String(),
	}

	return

	//return nil
}
