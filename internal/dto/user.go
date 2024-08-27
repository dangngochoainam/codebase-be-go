package dto

import "example/entity"

// Request - Response UseCase
type FindUsersRequestDTO struct {
	Username string
}
type FindUsersResponseDTO struct {
	Users []*entity.User
}

// Input - Output Repository
type FindUsersInput struct {
	Username string
}
