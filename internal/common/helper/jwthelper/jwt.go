package jwthelper

import (
	"errors"
	"example/internal/common/helper/loghelper"
	"example/internal/common/helper/structhelper"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type GenerateTokenOutput struct {
	AccessToken  string
	RefreshToken string
}
type TokenPayloadPublic struct {
	Key string `json:"key"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
}

type JwtConfigOptions struct {
	JwtSecret string
}

type (
	JwtHelper interface {
		CreateToken(payload *TokenPayloadPublic, tokenExpire int64) (string, error)
		VerifyToken(tokenString string) (*TokenPayloadPublic, error)
		GenerateToken(payload *TokenPayloadPublic, accessTokenLifeTime int64, refreshTokenLifeTime int64) (*GenerateTokenOutput, error)
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

func (j *jwtHelper) CreateToken(payload *TokenPayloadPublic, tokenExpire int64) (string, error) {
	if payload.Exp == 0 || payload.Iat == 0 || tokenExpire != 0 {
		currentTime := time.Now()
		iat := currentTime.Unix()
		exp := currentTime.Add(time.Second * time.Duration(tokenExpire)).Unix()
		payload.Exp = exp
		payload.Iat = iat
	}

	claims := &jwt.MapClaims{
		"key": payload.Key,
		"iat": payload.Iat,
		"exp": payload.Exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.options.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtHelper) GenerateToken(payload *TokenPayloadPublic, accessTokenLifeTime int64, refreshTokenLifeTime int64) (*GenerateTokenOutput, error) {
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

func (j *jwtHelper) VerifyToken(tokenString string) (*TokenPayloadPublic, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.options.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		loghelper.Logger.Error("invalid token")
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		loghelper.Logger.Error("Got error while convert token claims")
		return nil, errors.New("invalid token")
	}
	tokenPayloadPublic := &TokenPayloadPublic{}
	err = structhelper.MapToStruct(tokenPayloadPublic, claims)
	if err != nil {
		loghelper.Logger.Error("Got error while convert map to struct")
		return nil, err
	}
	return tokenPayloadPublic, nil
}
