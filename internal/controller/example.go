package controller

import (
	"example/internal/common/helper/jwthelper"
	"example/internal/common/helper/loghelper"
	"example/internal/common/helper/redishelper"
	"example/internal/common/helper/responsehelper"
	"example/internal/dto"
	"example/internal/usecase"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	ExampleController interface {
		JwtTest(ctx *gin.Context)
		JwtVerifyTest(ctx *gin.Context)
		GoroutineTest(ctx *gin.Context)
		RedisTest(ctx *gin.Context)
		MutexTest(ctx *gin.Context)
		ValidateTest(ctx *gin.Context)
	}

	exampleController struct {
		exampleUseCase usecase.ExampleUseCase
		redisSession   redishelper.RedisSessionHelper
		jwtHelper      jwthelper.JwtHelper
	}
)

func NewExampleController(exampleUseCase usecase.ExampleUseCase, redisSession redishelper.RedisSessionHelper, jwtHelper jwthelper.JwtHelper) ExampleController {
	return &exampleController{
		exampleUseCase: exampleUseCase,
		redisSession:   redisSession,
		jwtHelper:      jwtHelper,
	}
}

func (u *exampleController) RedisTest(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}
	token := "token is here"
	err := u.redisSession.Set(appC.C.Request.Context(), "token", token, time.Second*30)
	if err != nil {
		loghelper.Logger.Error("Got error while redis set, err: %v", err)
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}
	appC.Response(http.StatusOK, responsehelper.SUCCESS, nil)
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
func (u *exampleController) ValidateTest(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	body := &dto.ValidateExampleRequestDTO{}
	err := appC.C.ShouldBindBodyWithJSON(body)
	if err != nil {
		loghelper.Logger.Errorf("Got error while binding body, err: %v", err)
		appC.Response(http.StatusBadRequest, responsehelper.INVALID_PARAMS, nil)
		return
	}
	appC.Response(200, responsehelper.SUCCESS, body)
}

func (u *exampleController) JwtTest(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	tokenPayload := &jwthelper.TokenPayloadPublic{
		Key: "aslkdfjadsfhlk",
	}

	token, err := u.jwtHelper.CreateToken(tokenPayload, 120)
	if err != nil {
		loghelper.Logger.Errorf("Got error while generating token, err: %v", err)
		appC.Response(http.StatusBadRequest, responsehelper.INVALID_PARAMS, nil)
		return
	}

	loghelper.Logger.Info("token: ", token)

	appC.Response(200, responsehelper.SUCCESS, token)
}

func (u *exampleController) JwtVerifyTest(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}

	accessToken := appC.C.GetHeader("Authorization")
	accessToken = accessToken[len("Bearer "):]

	tokenPayloadPublic, err := u.jwtHelper.VerifyToken(accessToken)
	if err != nil {
		loghelper.Logger.Errorf("Got error while verifing token, err: %v", err)
		appC.Response(http.StatusBadRequest, responsehelper.INVALID_PARAMS, nil)
		return
	}

	appC.Response(200, responsehelper.SUCCESS, tokenPayloadPublic)
}
