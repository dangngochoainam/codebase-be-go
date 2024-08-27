package router

import (
	"example/internal/controller"
	"example/internal/diregistry"

	"github.com/gin-gonic/gin"
)

func registerUserRouter(router *gin.RouterGroup) {
	userController := diregistry.GetDependency(diregistry.UserControllerDIName).(controller.UserController)

	router.GET("/", userController.GetUserList)
}
