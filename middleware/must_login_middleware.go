package middleware

import (
	"github.com/gin-gonic/gin"
	"login-management-go/service"
)

type MustLoginMiddleware struct {
	service.SessionService
}

func NewMustLoginMiddleware(sessionService service.SessionService) UserMiddleware {
	return &MustLoginMiddleware{SessionService: sessionService}
}

func (m *MustLoginMiddleware) Before(ctx *gin.Context) {
	currentUser := m.SessionService.CurrentUser(ctx.Request)
	if currentUser.ID == 0 {
		ctx.Redirect(302, "/users/login")
	}
}
