package router

import (
	"example/internal/controller"
	"example/internal/diregistry"
	"example/internal/dto"
	"example/internal/middleware"

	"github.com/gin-gonic/gin"
)

func registerUserRouter(router *gin.RouterGroup, middleware middleware.Middleware) {
	userController := diregistry.GetDependency(diregistry.UserControllerDIName).(controller.UserController)

	router.POST("/", middleware.ValidateRequestMiddleware(&dto.CreateUserRequestDTO{}), userController.CreateUser)
	router.POST("/list", userController.CreateManyUser)
	router.GET("/one", userController.FindOneUser)
	router.GET("/", userController.FindUsers)
	router.PUT("/:id", userController.UpdateUserById)
	router.PUT("/", userController.UpdateUser)
	router.DELETE("/:id", userController.SoftDeleteUser)

}
