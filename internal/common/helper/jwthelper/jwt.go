package jwthelper

import (
	"errors"
	"example/internal/common/helper/loghelper"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type GenerateTokenOutput struct {
	AccessToken  string
	RefreshToken string
}
type TokenPayloadInput struct {
	Id string
}
type JwtConfigOptions struct {
	JwtSecret string
}

type (
	JwtHelper interface {
		CreateToken(payload *TokenPayloadInput, tokenExpire int64) (string, error)
		VerifyToken(token string) error
		GenerateToken(payload *TokenPayloadInput, accessTokenLifeTime int64, refreshTokenLifeTime int64) (*GenerateTokenOutput, error)
	}

	jwtHelper struct {
		options *JwtConfigOptions
	}
)

func NewJwt(options *JwtConfigOptions) JwtHelper {
	return &jwtHelper{
		options: options,
	}
}

func (j *jwtHelper) CreateToken(payload *TokenPayloadInput, tokenExpire int64) (string, error) {
	currentTime := time.Now()
	iat := currentTime.Unix()
	exp := currentTime.Add(time.Second * time.Duration(tokenExpire)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  payload.Id,
			"iat": iat,
			"exp": exp,
		})

	tokenString, err := token.SignedString([]byte(j.options.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtHelper) GenerateToken(payload *TokenPayloadInput, accessTokenLifeTime int64, refreshTokenLifeTime int64) (*GenerateTokenOutput, error) {
	accessToken, err := j.CreateToken(payload, accessTokenLifeTime)
	if err != nil {
		return nil, err
	}
	refreshToken, err := j.CreateToken(payload, refreshTokenLifeTime)
	if err != nil {
		return nil, err
	}

	result := &GenerateTokenOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return result, nil
}

func (j *jwtHelper) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.options.JwtSecret), nil
	})

	if err != nil {
		return err
	}
	if !token.Valid {
		loghelper.Logger.Error("invalid token")
		return errors.New("invalid token")
	}

	return nil
}
