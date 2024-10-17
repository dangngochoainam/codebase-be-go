package controller

import (
	"example/internal/common/helper/responsehelper"
	"example/internal/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	ExampleController interface {
		GoroutineTest(ctx *gin.Context)
		MutexTest(ctx *gin.Context)
	}

	exampleController struct {
		exampleUseCase usecase.ExampleUseCase
	}
)

func NewExampleController(exampleUseCase usecase.ExampleUseCase) ExampleController {
	return &exampleController{
		exampleUseCase: exampleUseCase,
	}
}

func (u *exampleController) GoroutineTest(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}
	data, err := u.exampleUseCase.GoroutineTest()
	if err != nil {
		fmt.Println("Error !!!")
		fmt.Println(err)
	}
	fmt.Println("Data: ", data)
	appC.Response(200, responsehelper.SUCCESS, data)
}

func (u *exampleController) MutexTest(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}
	go u.exampleUseCase.MutexTest()
	appC.Response(200, responsehelper.SUCCESS, "Ok")
}
