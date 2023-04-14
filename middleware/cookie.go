package middleware

import (
	"github.com/gin-gonic/gin"
	"login-management-go/auth"
	"login-management-go/helper"
	"login-management-go/model/web"
	"net/http"
)

func DefaultCookie(ctx *gin.Context) {
	jwtAuth := auth.NewJWTAuth()
	token, err := jwtAuth.GenerateToken(web.JWTClaims{SID: "asas"})
	helper.PanicIfError(err)

	c := new(http.Cookie)
	c.Name = "i-bell-a"
	c.Value = token
	c.Path = "/"
	ctx.Request.AddCookie(c)
}
