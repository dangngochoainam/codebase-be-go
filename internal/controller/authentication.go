package controller

import (
	"example/internal/common/helper/commonhelper"
	"example/internal/common/helper/jwthelper"
	"example/internal/common/helper/loghelper"
	"example/internal/common/helper/redishelper"
	"example/internal/common/helper/responsehelper"
	"example/internal/dto"
	"example/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	AuthenticationController interface {
		Login(ctx *gin.Context)
		Logout(ctx *gin.Context)
		RefreshToken(ctx *gin.Context)
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

func (a *authenticationController) Logout(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	key, ok := appC.C.Get(string(commonhelper.ContextKey_Key))
	if !ok {
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	input := &dto.LogoutRequestDTO{
		Key: key.(string),
	}
	data, err := a.authenticationUseCase.Logout(input)
	if err != nil {
		loghelper.Logger.Errorf("Got error while logouting user, err: %v", err)
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	appC.Response(200, responsehelper.SUCCESS, data)
}

func (a *authenticationController) RefreshToken(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	input := &dto.RefreshRequestDTO{}
	err := appC.C.ShouldBindBodyWithJSON(input)
	if err != nil {
		loghelper.Logger.Errorf("Got error while binding body, err: %v", err)
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	data, err := a.authenticationUseCase.RefreshToken(input)
	if err != nil {
		loghelper.Logger.Errorf("Got error while refreshing user, err: %v", err)
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(200, responsehelper.SUCCESS, data)
}
