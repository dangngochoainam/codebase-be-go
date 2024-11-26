package router

import (
	"example/internal/controller"
	"example/internal/diregistry"
	"example/internal/middleware"

	"github.com/gin-gonic/gin"
)

func registerAccountRouter(router *gin.RouterGroup, middleware middleware.Middleware) {
	accountController := diregistry.GetDependency(diregistry.AccountControllerDIName).(controller.AccountController)

	router.POST("/fun-transfer", accountController.TransferMoney)

}
