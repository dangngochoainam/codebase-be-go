package controller

import (
	"example/internal/common/helper/jwthelper"
	"example/internal/common/helper/loghelper"
	"example/internal/common/helper/redishelper"
	"example/internal/common/helper/responsehelper"
	"example/internal/dto"
	"example/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	AuthenticationController interface {
		Login(ctx *gin.Context)
	}

	authenticationController struct {
		authenticationUseCase usecase.AuthenticationUseCase
		redisSession          redishelper.RedisSessionHelper
		jwtHelper             jwthelper.JwtHelper
	}
)

func NewAuthenticationController(authUseCase usecase.AuthenticationUseCase, redisSession redishelper.RedisSessionHelper, jwtHelper jwthelper.JwtHelper) AuthenticationController {
	return &authenticationController{
		authenticationUseCase: authUseCase,
		redisSession:          redisSession,
		jwtHelper:             jwtHelper,
	}
}

func (a *authenticationController) Login(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	input := &dto.LoginRequestDTO{}
	err := appC.C.ShouldBindBodyWithJSON(input)
	if err != nil {
		loghelper.Logger.Errorf("Got error while binding body, err: %v", err)
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	data, err := a.authenticationUseCase.Authentication(input)
	if err != nil {
		loghelper.Logger.Errorf("Got error while authentication user, err: %v", err)
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(200, responsehelper.SUCCESS, data)
}
