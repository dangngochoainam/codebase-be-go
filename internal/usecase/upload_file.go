package usecase

import (
	"example/internal/common/helper/loghelper"
	"example/internal/common/utils"
	"fmt"
	"mime/multipart"
)

type (
	UploadFile interface {
		UploadSingleFile(file *multipart.FileHeader, fileName string) (bool, error)
	}

	uploadFile struct {
	}
)

func NewUploadFile() UploadFile {
	return &uploadFile{}
}

func (u *uploadFile) UploadSingleFile(file *multipart.FileHeader, fileName string) (bool, error) {
	dst := fmt.Sprintf("%s/%s", "./internal/static", fileName)
	err := utils.SaveUploadedFile(file, dst)
	if err != nil {
		loghelper.Logger.Errorf("Got error while saving file, err: %s", err)
		return false, err
	}
	return true, nil
}
