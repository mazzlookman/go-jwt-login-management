package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultCookie(ctx *gin.Context) {
	c := new(http.Cookie)
	c.Name = "i-bell-a"
	c.Value = "none"
	c.Path = "/"
	ctx.Request.AddCookie(c)
}
