package router

import (
	"example/internal/controller"
	"example/internal/diregistry"

	"github.com/gin-gonic/gin"
)

func registerUserRouter(router *gin.RouterGroup) {
	userController := diregistry.GetDependency(diregistry.UserControllerDIName).(controller.UserController)

	router.POST("/", userController.CreateUser)
	router.POST("/list", userController.CreateManyUser)
	router.GET("/one", userController.FindOneUser)
	router.GET("/", userController.FindUsers)
	router.PUT("/:id", userController.UpdateUserById)
	router.PUT("/", userController.UpdateUser)
	router.DELETE("/:id", userController.SoftDeleteUser)

}
