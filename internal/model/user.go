package model

//type UserCreateInput struct {
//	Name     string
//	Password string
//	Email    string
//}

type UserSignInInput struct {
	Username  string
	Password  string
	SubUserId int
}
