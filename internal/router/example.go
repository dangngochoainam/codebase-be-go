package router

import (
	"example/internal/controller"
	"example/internal/diregistry"
	"example/internal/dto"
	"example/internal/middleware"

	"github.com/gin-gonic/gin"
)

func registerExampleRouter(router *gin.RouterGroup) {
	exampleController := diregistry.GetDependency(diregistry.ExampleControllerDIName).(controller.ExampleController)

	router.GET("/goroutine-test", exampleController.GoroutineTest)
	router.GET("/mutex-test", exampleController.MutexTest)
	router.POST("/validate-test", middleware.ValidateRequestMiddleware(&dto.ValidateExam	pleRequestDTO{}), exampleController.ValidateTest)
}
