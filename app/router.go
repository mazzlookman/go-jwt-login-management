package app

import (
	"github.com/gin-gonic/gin"
	"login-management-go/auth"
	"login-management-go/controller"
	"login-management-go/middleware"
	"login-management-go/repository"
	"login-management-go/service"
)

func NewRouter() *gin.Engine {
	db := DBConnection()
	userRepository := repository.NewUserRepository(db)
	sessionRepository := repository.NewSessionRepository(db)
	userService := service.NewUserService(userRepository)
	jwtAuth := auth.NewJWTAuth()
	sessionService := service.NewSessionService(userRepository, sessionRepository, jwtAuth)
	userController := controller.NewUserController(userService, sessionService, jwtAuth)
	homeController := controller.NewHomeController(sessionService)
	mustLoginMiddleware := middleware.NewMustLoginMiddleware(sessionService)
	mustNotLoginMiddleware := middleware.NewMustNotLoginMiddleware(sessionService)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.gohtml")
	router.Use(middleware.DefaultCookie)

	router.GET("/", mustLoginMiddleware.Before, homeController.Home)
	router.GET("/users/register", mustNotLoginMiddleware.Before, userController.Register)
	router.POST("/users/register", userController.PostRegister)
	router.GET("/users/login", mustNotLoginMiddleware.Before, userController.Login)
	router.POST("/users/login", userController.PostLogin)
	router.GET("/users/logout", mustLoginMiddleware.Before, userController.Logout)

	return router
}
