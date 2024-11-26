package router

import (
	"example/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(middleware middleware.Middleware) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Gin!",
		})
	})

	router.Use(middleware.SetTimeMsMiddleware())
	router.Use(middleware.SetTraceIdMiddleware())
	router.Use(middleware.AuthenticationMiddleware())

	// router.POST("/auth", api.GetAuth)
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.POST("/upload", api.UploadImage)

	apiV1 := router.Group("/api/v1")
	registerUploadFileRouter(apiV1.Group("/upload-file"), middleware)
	registerAuthenticationRouter(apiV1.Group("/auth"), middleware)
	registerUserRouter(apiV1.Group("/users"), middleware)
	registerAccountRouter(apiV1.Group("/accounts"), middleware)
	registerExampleRouter(apiV1.Group("/example"), middleware)

	return router
}
