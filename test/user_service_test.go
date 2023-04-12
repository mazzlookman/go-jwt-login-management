package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"login-management-go/app"
	"login-management-go/model/web"
	"login-management-go/repository"
	"login-management-go/service"
	"testing"
)

var userService = service.NewUserService(repository.NewUserRepository(app.DBConnection()))

func TestRegister(t *testing.T) {
	input := web.UserRegisterInput{
		Name:     "Mario",
		Email:    "mario@test.com",
		Password: "123",
	}

	user, err := userService.Register(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Equal(t, "mario@test.com", user.Email)
	fmt.Println(user)
}

func TestRegisterDuplicate(t *testing.T) {
	input := web.UserRegisterInput{
		Name:     "Otong",
		Email:    "otong@test.com",
		Password: "123",
	}

	_, err := userService.Register(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Error(t, err, "Email not available")
}

func TestLogin(t *testing.T) {
	input := web.UserLoginInput{
		Email:    "otong@test.com",
		Password: "123",
	}

	user, err := userService.Login(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Equal(t, "otong@test.com", user.Email)
}

func TestLoginWrongEmail(t *testing.T) {
	input := web.UserLoginInput{
		Email:    "otong2@test.com",
		Password: "123",
	}

	user, err := userService.Login(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Error(t, err, "Login failed")
	assert.Equal(t, 0, user.ID)
}

func TestLoginWrongPassword(t *testing.T) {
	input := web.UserLoginInput{
		Email:    "otong@test.com",
		Password: "1234",
	}

	_, err := userService.Login(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Error(t, err, "Login failed")
}
