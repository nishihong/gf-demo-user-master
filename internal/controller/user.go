package controller

import (
	"context"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"

	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

var User = cUser{}

type cUser struct{}

//func (c *cUser) SignUp(ctx context.Context, req *v1.UserSignUpReq) (res *v1.UserSignUpRes, err error) {
//	err = service.User().Create(ctx, model.UserCreateInput{
//		Passport: req.Passport,
//		Password: req.Password,
//		Nickname: req.Nickname,
//	})
//	return
//}
func (c *cUser) SignIn(ctx context.Context, req *v1.UserSignInReq) (res *v1.UserSignInRes, err error) {
	Result, err := service.User().SignIn(ctx, model.UserSignInInput{
		Username:  req.Username,
		Password:  req.Password,
		SubUserId: req.SubUserId,
	})

	if err != nil {
		return nil, err
	}

	res = &v1.UserSignInRes{
		AccessToken: Result["AccessToken"].(string),
		TokenType:   Result["TokenType"].(string),
		ExpiresIn:   Result["ExpiresIn"].(int),
	}

	return
}

func (c *cUser) IsSignedIn(ctx context.Context, req *v1.UserIsSignedInReq) (res *v1.UserIsSignedInRes, err error) {
	res = &v1.UserIsSignedInRes{
		IsOnline: service.User().IsSignedIn(ctx),
	}

	return
}

func (c *cUser) SignOut(ctx context.Context, req *v1.UserSignOutReq) (res *v1.UserSignOutRes, err error) {
	err = service.User().SignOut(ctx)
	return
}

//func (c *cUser) CheckPassport(ctx context.Context, req *v1.UserCheckPassportReq) (res *v1.UserCheckPassportRes, err error) {
//	available, err := service.User().IsPassportAvailable(ctx, req.Passport)
//	if err != nil {
//		return nil, err
//	}
//	if !available {
//		return nil, gerror.Newf(`Passport "%s" is already token by others`, req.Passport)
//	}
//	return
//}

//func (c *cUser) CheckNickName(ctx context.Context, req *v1.UserCheckNickNameReq) (res *v1.UserCheckNickNameRes, err error) {
//	available, err := service.User().IsNicknameAvailable(ctx, req.Nickname)
//	if err != nil {
//		return nil, err
//	}
//	if !available {
//		return nil, gerror.Newf(`Nickname "%s" is already token by others`, req.Nickname)
//	}
//	return
//}
//
//func (c *cUser) Profile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
//	res = &v1.UserProfileRes{
//		User: service.User().GetProfile(ctx),
//	}
//	return
//}
