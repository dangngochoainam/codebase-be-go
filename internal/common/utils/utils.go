package utils

import (
	"errors"
	"example/internal/common/helper/loghelper"
	"golang.org/x/crypto/bcrypt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		loghelper.Logger.Errorf("Got error while hashing password, err: %v", err)
		return "", errors.New("error hashing password")
	}
	return string(bytes), nil
}

func VerifyPassword(rawPassword string, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword))
	if err != nil {
		loghelper.Logger.Errorf("Got error while verifing password, err: %v", err)
		return errors.New("Password not found!!")
	}
	return nil
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
