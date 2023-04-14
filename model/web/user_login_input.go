package web

type UserLoginInput struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
