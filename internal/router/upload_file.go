package router

import (
	"example/internal/controller"
	"example/internal/diregistry"
	"example/internal/middleware"

	"github.com/gin-gonic/gin"
)

func registerUploadFileRouter(router *gin.RouterGroup, middleware middleware.Middleware) {
	uploadFileController := diregistry.GetDependency(diregistry.UploadFileControllerDIName).(controller.UploadFileController)
	
	router.POST("/", uploadFileController.UploadSingleFile)
}
