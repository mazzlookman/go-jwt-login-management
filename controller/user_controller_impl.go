package controller

import (
	"github.com/gin-gonic/gin"
	"login-management-go/helper"
	"login-management-go/model/web"
	"login-management-go/service"
	"net/http"
)

type UserControllerImpl struct {
	service.UserService
	service.SessionService
}

func (c *UserControllerImpl) Logout(ctx *gin.Context) {
	c.SessionService.Destroy(ctx.Writer, ctx.Request)
	currentUser := c.SessionService.CurrentUser(ctx.Request)
	ctx.HTML(200, "home.gohtml", currentUser)
}

func (c *UserControllerImpl) Register(ctx *gin.Context) {
	ctx.HTML(200, "register.gohtml", nil)
}

func (c *UserControllerImpl) PostRegister(ctx *gin.Context) {
	input := web.UserRegisterInput{}
	err := ctx.ShouldBind(&input)
	helper.PanicIfError(err)

	_, err = c.UserService.Register(input)
	helper.PanicIfError(err)

	ctx.Redirect(http.StatusFound, "/users/login")
}

func (c *UserControllerImpl) Login(ctx *gin.Context) {
	ctx.HTML(200, "login.gohtml", nil)
}

func (c *UserControllerImpl) PostLogin(ctx *gin.Context) {
	input := web.UserLoginInput{}
	err := ctx.ShouldBind(&input)
	helper.PanicIfError(err)

	user, err := c.UserService.Login(input)
	helper.PanicIfError(err)

	c.SessionService.Create(ctx.Writer, user.ID)

	ctx.Redirect(http.StatusFound, "/")
}

func NewUserController(userService service.UserService, sessionService service.SessionService) UserController {
	return &UserControllerImpl{UserService: userService, SessionService: sessionService}
}
