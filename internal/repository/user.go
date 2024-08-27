package repository

import (
	"example/config"
	"example/entity"
	"example/internal/dto"
)

type (
	UserRepository interface {
		FindUsers(input *dto.FindUsersInput) ([]*entity.User, error)
	}

	userRepository struct {
		cfg *config.Config
	}
)

func NewUserRepository(cfg *config.Config) UserRepository {
	return &userRepository{
		cfg: cfg,
	}
}

func (u *userRepository) FindUsers(input *dto.FindUsersInput) ([]*entity.User, error) {
	result := []*entity.User{}
	user := &entity.User{
		Username: input.Username,
	}
	result = append(result, user)
	return result, nil
}
