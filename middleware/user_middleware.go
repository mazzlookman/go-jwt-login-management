package middleware

import (
	"github.com/gin-gonic/gin"
)

type UserMiddleware interface {
	Before(ctx *gin.Context)
}
