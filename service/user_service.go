package service

import (
	"login-management-go/model/domain"
	"login-management-go/model/web"
)

type UserService interface {
	Register(input web.UserRegisterInput) (domain.User, error)
	Login(input web.UserLoginInput) (domain.User, error)
	UpdateProfile(input web.UserUpdateProfileInput) (domain.User, error)
}
