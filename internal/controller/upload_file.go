package controller

import (
	"example/internal/common/helper/loghelper"
	"example/internal/common/helper/responsehelper"
	"example/internal/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type (
	UploadFileController interface {
		UploadSingleFile(ctx *gin.Context)
	}

	uploadFileController struct {
		uploadFile usecase.UploadFile
	}
)

func NewUploadFileController(uploadFile usecase.UploadFile) UploadFileController {
	return &uploadFileController{
		uploadFile: uploadFile,
	}
}

func (uc *uploadFileController) UploadSingleFile(ctx *gin.Context) {
	appC := responsehelper.Gin{
		C: ctx,
	}
	// handle request from client
	file, err := appC.C.FormFile("file")
	if err != nil {
		loghelper.Logger.Errorf("Got error while getting file, err: %v", err)
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	// call usecase handle business
	filename := fmt.Sprintf("%s-%s", uuid.New().String(), file.Filename)
	data, err := uc.uploadFile.UploadSingleFile(file, filename)
	// handle response into client
	if err != nil {
		loghelper.Logger.Errorf("Got error while creating user, err: %v", err)
		appC.Response(http.StatusBadRequest, responsehelper.ERROR, nil)
		return
	}

	appC.Response(http.StatusOK, responsehelper.SUCCESS, data)
	return
}
