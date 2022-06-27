package controller

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/service"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"
)

var SdkUserSerial = cSdkUserSerial{}

type cSdkUserSerial struct{}

func (a *cSdkUserSerial) GetUserSerialList(ctx context.Context, req *v1.GetUserSerialListReq) (res *v1.GetUserSerialListRes, err error) {

	res = &v1.GetUserSerialListRes{
		Data: service.SdkUserSerial().GetUserSerialList(ctx, req.Id),
	}

	return
}

func (a *cSdkUserSerial) GetKey(ctx context.Context, req *v1.GetKeyReq) (res *v1.GetKeyRes, err error) {

	res = &v1.GetKeyRes{
		Data: service.SdkUserSerial().GetKey(ctx, req.Id),
	}

	return
}

func (a *cSdkUserSerial) EditUserSerial(ctx context.Context, req *v1.EditUserSerialReq) (res *v1.EditUserSerialRes, err error) {

	err = service.SdkUserSerial().EditUserSerial(ctx, req.Id, req.Name)

	if err != nil {
		return nil, err
	}

	return
}

func (a *cSdkUserSerial) CreateUserSerial(ctx context.Context, req *v1.CreateUserSerialReq) (res *v1.CreateUserSerialRes, err error) {

	err = service.SdkUserSerial().CreateUserSerial(ctx, req.SdkUserProductId, req.Name)

	if err != nil {
		return nil, err
	}

	return
}

func (a *cSdkUserSerial) DeleteUserSerial(ctx context.Context, req *v1.DeleteUserSerialReq) (res *v1.DeleteUserSerialRes, err error) {

	err = service.SdkUserSerial().DeleteUserSerial(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return
}
