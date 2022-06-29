package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

//type UserProfileReq struct {
//	g.Meta `path:"/user/profile" method:"get" tags:"UserService" summary:"Get the profile of current user"`
//}
//type UserProfileRes struct {
//	*entity.User
//}

//type UserSignUpReq struct {
//	g.Meta    `path:"/user/sign-up" method:"post" tags:"UserService" summary:"Sign up a new user account"`
//	Passport  string `v:"required|length:6,16"`
//	Password  string `v:"required|length:6,16"`
//	Password2 string `v:"required|length:6,16|same:Password"`
//	Nickname  string
//}
//type UserSignUpRes struct{}

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

//type UserCheckPassportReq struct {
//	g.Meta   `path:"/user/check-passport" method:"post" tags:"UserService" summary:"Check passport available"`
//	Passport string `v:"required"`
//}
//type UserCheckPassportRes struct{}

//type UserCheckNickNameReq struct {
//	g.Meta   `path:"/user/check-passport" method:"post" tags:"UserService" summary:"Check nickname available"`
//	Nickname string `v:"required"`
//}
//type UserCheckNickNameRes struct{}

type UserIsSignedInReq struct {
	g.Meta `path:"/user/is-signed-in" method:"post" tags:"UserService" summary:"Check current user is already signed-in"`
}
type UserIsSignedInRes struct {
	OK bool `dc:"True if current user is signed in; or else false"`
}

type UserSignOutReq struct {
	g.Meta `path:"/user/sign-out" method:"post" tags:"UserService" summary:"Sign out current user"`
}
type UserSignOutRes struct{}
