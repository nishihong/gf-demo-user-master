package controller

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/model"

	"github.com/gogf/gf-demo-user/v2/internal/service"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"
)

var SdkRules = cSdkRules{}

type cSdkRules struct{}

func (a *cSdkRules) GetRuleTypeList(ctx context.Context, req *v1.GetRuleTypeReq) (res *v1.GetRuleTypeRes, err error) {

	res = &v1.GetRuleTypeRes{
		Data: map[int]string{
			1: "PC端",
			2: "移动端",
			3: "PC端+移动端",
		},
	}

	return
}

func (a *cSdkRules) GetUserIpList(ctx context.Context, req *v1.GetUserIpReq) (res *v1.GetUserIpRes, err error) {

	pc := []string{"127.0.0.1", "127.0.0.5", "127.0.0.6", "127.0.0.7", "127.0.0.8", "127.0.0.9"}
	mobile := []string{"127.0.0.1"}
	res = &v1.GetUserIpRes{
		Pc:     pc,
		Mobile: mobile,
	}

	return
}

func (a *cSdkRules) GetSourceServersNum(ctx context.Context, req *v1.GetSourceServersNumReq) (res *v1.GetSourceServersNumRes, err error) {

	Result, err := service.SdkRules().GetSourceServersNum(ctx, req.UserSerialId)

	if err != nil {
		return nil, err
	}

	res = &v1.GetSourceServersNumRes{
		Data: Result,
	}

	return
}

func (a *cSdkRules) GetSdkRuleList(ctx context.Context, req *v1.GetSdkRuleListReq) (res *v1.GetSdkRuleListRes, err error) {

	Result, err := service.SdkRules().GetSdkRuleList(ctx, model.GetSdkRuleListInput{
		UserSerialId: req.UserSerialId,
		Type:         req.Type,
		Keyword:      req.Keyword,
		Port:         req.Port,
		Page:         req.Page,
		Limit:        req.Limit,
	})
	if err != nil {
		return nil, err
	}

	res = &v1.GetSdkRuleListRes{
		Data:        Result.Data,
		CurrentPage: Result.CurrentPage,
		Total:       Result.Total,
		PerPage:     Result.PerPage,
	}

	return
}

//func (a *cSdkRules) EditRule(ctx context.Context, req *v1.EditUserSerialReq) (res *v1.EditUserSerialRes, err error) {
//
//	err = service.SdkRules().EditRule(ctx, req.Id, req.Name)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return
//}
//
//func (a *cSdkRules) CreateRule(ctx context.Context, req *v1.CreateUserSerialReq) (res *v1.CreateUserSerialRes, err error) {
//
//	err = service.SdkRules().CreateRule(ctx, req.SdkUserProductId, req.Name)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return
//}

func (a *cSdkRules) DeleteRule(ctx context.Context, req *v1.DeleteRuleReq) (res *v1.DeleteRuleRes, err error) {

	err = service.SdkRules().DeleteRule(ctx, req.Ids, req.UserSerialId)
	if err != nil {
		return nil, err
	}

	return
}
