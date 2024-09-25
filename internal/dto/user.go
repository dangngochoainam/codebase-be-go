package dto

import "example/entity"

// Request - Response UseCase
type FindUsersRequestDTO struct {
	PagingRequestDTO
	Username string `form:"userName"`
}
type FindUsersResponseDTO struct {
	*PagingResponse
	List []*entity.User `json:"list"`
}

// Input - Output Repository
type FindUsersInput struct {
	Username string
}
