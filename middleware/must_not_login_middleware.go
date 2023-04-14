package middleware

import (
	"github.com/gin-gonic/gin"
	"login-management-go/service"
)

type MustNotLoginMiddleware struct {
	service.SessionService
}

func NewMustNotLoginMiddleware(sessionService service.SessionService) UserMiddleware {
	return &MustNotLoginMiddleware{SessionService: sessionService}
}

func (m *MustNotLoginMiddleware) Before(ctx *gin.Context) {
	currentUser := m.SessionService.CurrentUser(ctx.Request)
	if currentUser.ID != 0 {
		ctx.Redirect(302, "/")
	}
}
