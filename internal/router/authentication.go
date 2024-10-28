package router

import (
	"example/internal/controller"
	"example/internal/diregistry"
	"example/internal/dto"
	"example/internal/middleware"
	"github.com/gin-gonic/gin"
)

func registerAuthenticationRouter(router *gin.RouterGroup, middleware middleware.Middleware) {
	authenticationController := diregistry.GetDependency(diregistry.AuthenticationControllerDIName).(controller.AuthenticationController)
	router.POST("/login", middleware.ValidateRequestMiddleware(&dto.LoginRequestDTO{}), authenticationController.Login)
}
