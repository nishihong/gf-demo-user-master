package model

type UserCreateInput struct {
	Name     string
	Password string
	Email    string
}

type UserSignInInput struct {
	Name     string
	Password string
}
