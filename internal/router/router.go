package router

import (
	"example/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
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

	// router.POST("/auth", api.GetAuth)
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.POST("/upload", api.UploadImage)

	apiV1 := router.Group("/api/v1")
	registerUserRouter(apiV1.Group("/users"))
	registerExampleRouter(apiV1.Group("/example"))

	return router
}
