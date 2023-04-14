package controller

import (
	"github.com/gin-gonic/gin"
	"login-management-go/service"
)

type HomeController interface {
	Home(ctx *gin.Context)
}

type HomeControllerImpl struct {
	service.SessionService
}

func NewHomeController(sessionService service.SessionService) HomeController {
	return &HomeControllerImpl{SessionService: sessionService}
}

func (c *HomeControllerImpl) Home(ctx *gin.Context) {
	currentUser := c.SessionService.CurrentUser(ctx.Request)
	ctx.HTML(200, "home.gohtml", currentUser)
}
