package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(ctx *gin.Context)
	PostRegister(ctx *gin.Context)
	Login(ctx *gin.Context)
	PostLogin(ctx *gin.Context)
	Logout(ctx *gin.Context)
}
