package controller

import (
	"example/internal/common/helper/copyhepler"
	"example/internal/common/helper/responsehelper"
	"example/internal/dto"
	"example/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	AccountController interface {
		TransferMoney(ctx *gin.Context)
	}

	accountController struct {
		accountUseCase usecase.AccountUseCase
		modelConverter copyhepler.ModelConverter
	}
)

func NewAccountController(accountUseCase usecase.AccountUseCase,
	modelConverter copyhepler.ModelConverter) AccountController {
	return &accountController{
		accountUseCase: accountUseCase,
		modelConverter: modelConverter,
	}
}

func (a *accountController) TransferMoney(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	//handle request from client
	input := &dto.TransferMoneyRequestDTO{}
	err := appC.C.ShouldBindBodyWithJSON(input)
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	//call usecase handle business
	err = a.accountUseCase.TransferMoney(input)

	// handle response into client
	if err != nil {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(http.StatusOK, responsehelper.SUCCESS, "Ok")
	return
}
