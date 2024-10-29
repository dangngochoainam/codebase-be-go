package utils

import (
	"errors"
	"example/internal/common/helper/loghelper"
	"golang.org/x/crypto/bcrypt"
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
