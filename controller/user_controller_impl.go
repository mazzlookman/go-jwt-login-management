package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"login-management-go/auth"
	"login-management-go/helper"
	"login-management-go/model/web"
	"login-management-go/service"
	"net/http"
)

type UserControllerImpl struct {
	service.UserService
	service.SessionService
	auth.JWTAuth
}

func (c *UserControllerImpl) Logout(ctx *gin.Context) {
	c.SessionService.Destroy(ctx.Writer, ctx.Request)
	//currentUser := c.SessionService.CurrentUser(ctx.Request)
	ctx.Redirect(302, "/")
}

func (c *UserControllerImpl) Register(ctx *gin.Context) {
	ctx.HTML(200, "register.gohtml", nil)
}

func (c *UserControllerImpl) PostRegister(ctx *gin.Context) {
	input := web.UserRegisterInput{}
	err := ctx.ShouldBind(&input)
	helper.PanicIfError(err)

	user, err := c.UserService.Register(input)
	helper.PanicIfError(err)

	//generate JWT
	token, err := c.JWTAuth.GenerateToken(web.JWTClaims{
		SID: uuid.New().String(),
	})

	//create session
	c.SessionService.Create(ctx.Writer, token, user.ID)

	ctx.Redirect(http.StatusFound, "/")
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

	//generate JWT
	token, err := c.JWTAuth.GenerateToken(web.JWTClaims{
		SID: uuid.New().String(),
	})

	//create session
	c.SessionService.Create(ctx.Writer, token, user.ID)

	ctx.Redirect(http.StatusFound, "/")
}

func NewUserController(userService service.UserService, sessionService service.SessionService, jwtAuth auth.JWTAuth) UserController {
	return &UserControllerImpl{UserService: userService, SessionService: sessionService, JWTAuth: jwtAuth}
}
