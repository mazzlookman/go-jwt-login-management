package formatter

import (
	"login-management-go/model/domain"
	"login-management-go/model/web"
)

func UserResponseFormatter(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Token:    user.Token,
	}
}
