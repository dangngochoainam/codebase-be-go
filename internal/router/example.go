package router

import (
	"example/internal/controller"
	"example/internal/diregistry"

	"github.com/gin-gonic/gin"
)

func registerExampleRouter(router *gin.RouterGroup) {
	exampleController := diregistry.GetDependency(diregistry.ExampleControllerDIName).(controller.ExampleController)

	router.GET("/goroutine-test", exampleController.GoroutineTest)
}
