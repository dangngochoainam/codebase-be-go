package usecase

import (
	"example/entity"
	"example/internal/dto"
	"example/internal/repository"
)

type (
	UserUseCase interface {
		FindUsers(input *dto.FindUsersRequestDTO) (*dto.FindUsersResponseDTO, error)
		CreateUser(input *dto.CreateUserRequestDTO) (bool, error)
	}

	userUseCase struct {
		userRepository repository.UserRepository
	}
)

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) FindUsers(input *dto.FindUsersRequestDTO) (*dto.FindUsersResponseDTO, error) {
	users, err := u.userRepository.FindUsers(&dto.FindUsersInput{
		Username: input.Username,
	})
	if err != nil {
		return nil, err
	}

	response := &dto.PagingResponse{
		CurrentPage: 1,
		TotalPages:  2,
		TotalItems:  3,
	}
	result := &dto.FindUsersResponseDTO{
		PagingResponse: response,
		List:           users,
		// List: users,
	}
	return result, nil
}

func (u *userUseCase) CreateUser(input *dto.CreateUserRequestDTO) (bool, error) {
	userEntity := &entity.User{
		Username: input.Username,
		Password: input.Password,
	}
	_, err := u.userRepository.CreateUser(userEntity)
	if err != nil {
		return false, err
	}

	return true, nil
}
