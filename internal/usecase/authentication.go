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
	"fmt"
	"time"
)

type (
	AuthenticationUseCase interface {
		Authentication(input *dto.LoginRequestDTO) (*dto.LoginResponseDTO, error)
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

	tokenPayloadInput := &jwthelper.TokenPayloadInput{
		Id: fmt.Sprintf("encrypt-%s", user.Id),
	}
	tokens, err := a.jwt.GenerateToken(tokenPayloadInput, a.cfg.Jwt.TokenLifeTime, a.cfg.Jwt.RefreshTokenLifeTime)
	if err != nil {
		return nil, err
	}

	redisSessionAccessToken := redishelper.GenerateRedisSessionKey("accessToken", user.Id)
	err = a.redisSession.Set(context.Background(), redisSessionAccessToken, user.Id, time.Second*time.Duration(a.cfg.Jwt.TokenLifeTime))
	if err != nil {
		return nil, err
	}

	redisSessionRefreshToken := redishelper.GenerateRedisSessionKey("refreshToken", user.Id)
	err = a.redisSession.Set(context.Background(), redisSessionRefreshToken, user.Id, time.Second*time.Duration(a.cfg.Jwt.RefreshTokenLifeTime))
	if err != nil {
		return nil, err
	}

	result := &dto.LoginResponseDTO{}
	a.modelConverter.ToModel(result, tokens)

	return result, nil
}
