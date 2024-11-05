package router

import (
	"example/internal/controller"
	"example/internal/diregistry"
	"example/internal/dto"
	"example/internal/middleware"

	"github.com/gin-gonic/gin"
)

func registerExampleRouter(router *gin.RouterGroup, middleware middleware.Middleware) {
	exampleController := diregistry.GetDependency(diregistry.ExampleControllerDIName).(controller.ExampleController)

	router.GET("/redis-test", exampleController.RedisTest)
	router.GET("/fetch-test", exampleController.FetchClientGet)
	router.GET("/jwt-test", exampleController.JwtTest)
	router.POST("/jwt-test", exampleController.JwtVerifyTest)
	router.GET("/goroutine-test", exampleController.GoroutineTest)
	router.GET("/mutex-test", exampleController.MutexTest)
	router.POST("/validate-test", middleware.ValidateRequestMiddleware(&dto.ValidateExampleRequestDTO{}), exampleController.ValidateTest)
}
