package repository

import (
	"example/config"
	"example/entity"
	"example/internal/common/helper/sqlormhelper"
	"example/internal/dto"
)

type (
	UserRepository interface {
		FindUsers(input *dto.FindUsersInput) ([]*entity.User, error)
		CreateUser(entity *entity.User) (*entity.User, error)
	}

	userRepository struct {
		cfg           *config.Config
		postgresOrmDb sqlormhelper.SqlGormDatabase
	}
)

func NewUserRepository(cfg *config.Config, postgresOrmDb sqlormhelper.SqlGormDatabase) UserRepository {
	return &userRepository{
		cfg:           cfg,
		postgresOrmDb: postgresOrmDb,
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

func (u *userRepository) CreateUser(entity *entity.User) (*entity.User, error) {
	db := u.postgresOrmDb.Open()
	result := db.Create(entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}
