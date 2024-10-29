package repository

import (
	"example/config"
	"example/entity"
	"example/internal/common/helper/sqlormhelper"
	"example/internal/common/helper/structhelper"
	"example/internal/dto"
)

type (
	UserRepository interface {
		CreateUser(entity *entity.User) (*entity.User, error)
		CreateUsers(entity []*entity.User) ([]*entity.User, error)
		FindOneUser(params *dto.FindOneUserInput) (*entity.User, error)
		FindUserById(id string) (*entity.User, error)
		FindUsers(input *dto.FindUsersInput) ([]*entity.User, error)
		UpdateUserById(id string, dataToUpdate *dto.UpdateUserInput) (int64, error)
		UpdateUser(condition *dto.UpdateUserCondInput, dataToUpdate *dto.UpdateUserInput) (int64, error)
		SoftDeleteUser(id string) (int64, error)
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

func (u *userRepository) CreateUser(entity *entity.User) (*entity.User, error) {
	db := u.postgresOrmDb.Open()

	result := db.Create(entity)
	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

func (u *userRepository) CreateUsers(entities []*entity.User) ([]*entity.User, error) {
	db := u.postgresOrmDb.Open()

	result := db.Create(entities)
	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func (u *userRepository) FindOneUser(input *dto.FindOneUserInput) (*entity.User, error) {
	db := u.postgresOrmDb.Open()

	user := &entity.User{}
	userCond := &entity.User{
		Username: input.Username,
	}

	result := db.Where(userCond).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *userRepository) FindUserById(id string) (*entity.User, error) {
	db := u.postgresOrmDb.Open()

	user := &entity.User{}
	userCond := &entity.User{
		Id: id,
	}

	result := db.Where(userCond).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *userRepository) FindUsers(input *dto.FindUsersInput) ([]*entity.User, error) {
	db := u.postgresOrmDb.Open()

	users := []*entity.User{}
	userCond := &entity.User{
		Username: input.Username,
	}

	result := db.Where(userCond).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (u *userRepository) UpdateUserById(id string, dataToUpdate *dto.UpdateUserInput) (int64, error) {
	db := u.postgresOrmDb.Open()

	user := &entity.User{}
	updateUserCond := &entity.User{
		Id: id,
	}
	fieldsToUpdate := structhelper.GetFieldName(*dataToUpdate)
	result := db.Model(user).Where(updateUserCond).Select(fieldsToUpdate).Updates(dataToUpdate)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}

func (u *userRepository) UpdateUser(input *dto.UpdateUserCondInput, dataToUpdate *dto.UpdateUserInput) (int64, error) {
	db := u.postgresOrmDb.Open()

	user := &entity.User{}
	updateUserCond := &entity.User{
		Age: input.Age,
	}
	fieldsToUpdate := structhelper.GetFieldName(*dataToUpdate)
	result := db.Model(user).Where(updateUserCond).Select(fieldsToUpdate).Updates(dataToUpdate)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}
func (u *userRepository) SoftDeleteUser(id string) (int64, error) {
	db := u.postgresOrmDb.Open()

	deleteUserCond := &entity.User{
		Id: id,
	}
	result := db.Where(deleteUserCond).Delete(&entity.User{})
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}
