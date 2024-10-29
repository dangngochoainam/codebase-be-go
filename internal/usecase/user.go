package usecase

import (
	"errors"
	"example/entity"
	"example/internal/common/helper/copyhepler"
	"example/internal/common/utils"
	"example/internal/dto"
	"example/internal/repository"
)

type (
	UserUseCase interface {
		CreateUser(input *dto.CreateUserRequestDTO) (bool, error)
		CreateUsers(input *dto.CreateUsersRequestDTO) (bool, error)
		FindOneUser(input *dto.FindOneUserRequestDTO) (*dto.FindOneUserResponseDTO, error)
		FindUsers(input *dto.FindUsersRequestDTO) (*dto.FindUsersResponseDTO, error)
		UpdateUserById(id string, dataToUpdate *dto.UpdateUserRequestDTO) (bool, error)
		UpdateUser(dataToUpdate *dto.UpdateUserRequestDTO) (bool, error)
		SoftDeleteUser(id string) (bool, error)
	}

	userUseCase struct {
		userRepository repository.UserRepository
		modelConverter copyhepler.ModelConverter
	}
)

func NewUserUseCase(userRepository repository.UserRepository, modelConverter copyhepler.ModelConverter) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
		modelConverter: modelConverter,
	}
}

func (u *userUseCase) CreateUser(input *dto.CreateUserRequestDTO) (bool, error) {
	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return false, err
	}
	userEntity := &entity.User{
		Username: input.Username,
		Password: hashPassword,
		Email:    input.Email,
		Age:      input.Age,
	}
	_, err = u.userRepository.CreateUser(userEntity)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *userUseCase) CreateUsers(input *dto.CreateUsersRequestDTO) (bool, error) {
	userEntities := []*entity.User{}

	for _, user := range input.Users {
		userEntity := &entity.User{
			Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
			Age:      user.Age,
		}
		userEntities = append(userEntities, userEntity)
	}
	_, err := u.userRepository.CreateUsers(userEntities)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *userUseCase) FindOneUser(input *dto.FindOneUserRequestDTO) (*dto.FindOneUserResponseDTO, error) {
	user, err := u.userRepository.FindOneUser(&dto.FindOneUserInput{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	result := &dto.FindOneUserResponseDTO{}
	u.modelConverter.ToModel(result, user)

	return result, nil
}

func (u *userUseCase) FindUsers(query *dto.FindUsersRequestDTO) (*dto.FindUsersResponseDTO, error) {
	input := &dto.FindUsersInput{}
	u.modelConverter.ToModel(input, query)

	data, err := u.userRepository.FindUsers(input)
	if err != nil {
		return nil, err
	}

	users := []*dto.FindOneUserResponseDTO{}
	for _, item := range data {
		user := &dto.FindOneUserResponseDTO{}
		u.modelConverter.ToModel(user, item)
		users = append(users, user)
	}

	response := &dto.PagingResponse{
		CurrentPage: 1,
		TotalPages:  2,
		TotalItems:  3,
	}
	result := &dto.FindUsersResponseDTO{
		PagingResponse: response,
		List:           users,
	}

	return result, nil
}

func (u *userUseCase) UpdateUserById(id string, body *dto.UpdateUserRequestDTO) (bool, error) {
	dataToUpdate := &dto.UpdateUserInput{}
	u.modelConverter.ToModel(dataToUpdate, body)

	rowsAffected, err := u.userRepository.UpdateUserById(id, dataToUpdate)
	if err != nil {
		return false, err
	}
	if rowsAffected == 0 {
		return false, errors.New("Update not effect!!")
	}

	return true, nil
}

func (u *userUseCase) UpdateUser(body *dto.UpdateUserRequestDTO) (bool, error) {
	dataToUpdate := &dto.UpdateUserInput{}
	u.modelConverter.ToModel(dataToUpdate, body)
	updateUserCondInput := &dto.UpdateUserCondInput{
		Age: 23,
	}

	rowsAffected, err := u.userRepository.UpdateUser(updateUserCondInput, dataToUpdate)
	if err != nil {
		return false, err
	}
	if rowsAffected == 0 {
		return false, errors.New("Update not effect!!")
	}

	return true, nil
}

func (u *userUseCase) SoftDeleteUser(id string) (bool, error) {
	rowsAffected, err := u.userRepository.SoftDeleteUser(id)
	if err != nil {
		return false, err
	}
	if rowsAffected == 0 {
		return false, errors.New("Update not effect!!")
	}

	return true, nil
}
