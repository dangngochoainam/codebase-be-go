package usecase

import (
	"context"
	"example/config"
	"example/internal/common/helper/copyhepler"
	"example/internal/common/helper/jwthelper"
	"example/internal/common/helper/redishelper"
	"example/internal/common/utils"
	"example/internal/dto"
	"example/internal/repository"
	"time"

	"github.com/google/uuid"
)

type (
	AuthenticationUseCase interface {
		Authentication(input *dto.LoginRequestDTO) (*dto.LoginResponseDTO, error)
		ClearSession(key string) error
		CreateSession(key string, value any, expiration time.Duration) error
		GenerateSession(key string, value any) error
		Logout(input *dto.LogoutRequestDTO) (string, error)
		RefreshToken(input *dto.RefreshRequestDTO) (*dto.RefreshResponseDTO, error)
	}

	authenticationUseCase struct {
		userRepository repository.UserRepository
		modelConverter copyhepler.ModelConverter
		jwt            jwthelper.JwtHelper
		cfg            *config.Config
		redisSession   redishelper.RedisSessionHelper
	}
)

func NewAuthenticationUseCase(userRepository repository.UserRepository,
	modelConverter copyhepler.ModelConverter, jwt jwthelper.JwtHelper, cfg *config.Config, redisSession redishelper.RedisSessionHelper) AuthenticationUseCase {
	return &authenticationUseCase{
		userRepository: userRepository,
		modelConverter: modelConverter,
		jwt:            jwt,
		cfg:            cfg,
		redisSession:   redisSession,
	}
}

func (a *authenticationUseCase) CreateSession(key string, value any, expiration time.Duration) error {
	err := a.redisSession.Set(context.Background(), key, value, expiration)
	if err != nil {
		return err
	}
	return nil
}

func (a *authenticationUseCase) GenerateSession(key string, value any) error {
	redisSessionAccessTokenKey := redishelper.GenerateRedisSessionKey(redishelper.ACCESS_TOKEN, key)
	err := a.CreateSession(redisSessionAccessTokenKey, value, time.Second*time.Duration(a.cfg.Jwt.TokenLifeTime))
	if err != nil {
		return err
	}

	redisSessionRefreshTokenKey := redishelper.GenerateRedisSessionKey(redishelper.REFRESH_TOKEN, key)
	err = a.CreateSession(redisSessionRefreshTokenKey, value, time.Second*time.Duration(a.cfg.Jwt.RefreshTokenLifeTime))
	if err != nil {
		return err
	}
	return nil
}

func (a *authenticationUseCase) Authentication(input *dto.LoginRequestDTO) (*dto.LoginResponseDTO, error) {
	findOneUserInput := &dto.FindOneUserInput{}
	a.modelConverter.ToModel(findOneUserInput, input)
	user, err := a.userRepository.FindOneUser(findOneUserInput)
	if err != nil {
		return nil, err
	}
	err = utils.VerifyPassword(input.Password, user.Password)
	if err != nil {
		return nil, err
	}

	uuidKey := uuid.New().String()
	tokenPayloadPublic := &jwthelper.TokenPayloadPublic{
		Key: uuidKey,
	}
	tokens, err := a.jwt.GenerateToken(tokenPayloadPublic, a.cfg.Jwt.TokenLifeTime, a.cfg.Jwt.RefreshTokenLifeTime)
	if err != nil {
		return nil, err
	}

	err = a.GenerateSession(tokenPayloadPublic.Key, user.Id)
	if err != nil {
		return nil, err
	}

	result := &dto.LoginResponseDTO{}
	a.modelConverter.ToModel(result, tokens)

	return result, nil
}

func (a *authenticationUseCase) ClearSession(key string) error {
	redisSessionAccessTokenKey := redishelper.GenerateRedisSessionKey(redishelper.ACCESS_TOKEN, key)
	err := a.redisSession.Del(context.Background(), redisSessionAccessTokenKey)
	if err != nil {
		return err
	}

	redisSessionRefreshTokenKey := redishelper.GenerateRedisSessionKey(redishelper.REFRESH_TOKEN, key)
	err = a.redisSession.Del(context.Background(), redisSessionRefreshTokenKey)
	if err != nil {
		return err
	}

	return nil
}

func (a *authenticationUseCase) Logout(input *dto.LogoutRequestDTO) (string, error) {
	err := a.ClearSession(input.Key)
	if err != nil {
		return "", err
	}
	return "Ok", nil
}

func (a *authenticationUseCase) RefreshToken(input *dto.RefreshRequestDTO) (*dto.RefreshResponseDTO, error) {
	payloadVerified, err := a.jwt.VerifyToken(input.RefreshToken)
	if err != nil {
		return nil, err
	}

	var userId string
	redisSessionKey := redishelper.GenerateRedisSessionKey(redishelper.REFRESH_TOKEN, payloadVerified.Key)
	err = a.redisSession.Get(context.Background(), redisSessionKey, &userId)
	if err != nil {
		return nil, err
	}

	uuidKey := uuid.New().String()
	tokenPayloadPublic := &jwthelper.TokenPayloadPublic{
		Key: uuidKey,
	}
	tokens, err := a.jwt.GenerateToken(tokenPayloadPublic, a.cfg.Jwt.TokenLifeTime, a.cfg.Jwt.RefreshTokenLifeTime)
	if err != nil {
		return nil, err
	}

	result := &dto.RefreshResponseDTO{}
	a.modelConverter.ToModel(result, tokens)

	err = a.ClearSession(payloadVerified.Key)
	if err != nil {
		return nil, err
	}
	err = a.GenerateSession(tokenPayloadPublic.Key, userId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
