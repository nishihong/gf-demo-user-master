package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserSignInReq struct {
	g.Meta    `path:"/api/authorization/login" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	Username  string `json:"username" v:"required"`
	Password  string `json:"password" v:"required"`
	SubUserId int    `json:"sub_user_id"`
}
type UserSignInRes struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type UserIsSignedInReq struct {
	g.Meta `path:"/user/is-signed-in" method:"post" tags:"UserService" summary:"Check current user is already signed-in"`
}
type UserIsSignedInRes struct {
	IsOnline bool `dc:"True if current user is signed in; or else false"`
}

type UserSignOutReq struct {
	g.Meta `path:"/user/sign-out" method:"post" tags:"UserService" summary:"Sign out current user"`
}
type UserSignOutRes struct{}
