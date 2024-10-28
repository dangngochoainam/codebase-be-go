package dto

// Request - Response UseCase
type CreateUserRequestDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required"`
}

type CreateUsersRequestDTO struct {
	Users []*CreateUserRequestDTO `json:"users"`
}

type FindOneUserRequestDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type FindOneUserResponseDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type FindUsersRequestDTO struct {
	PagingRequestDTO
	Username string `form:"userName"`
	Age      string `form:"age"`
}
type FindUsersResponseDTO struct {
	*PagingResponse
	List []*FindOneUserResponseDTO `json:"list"`
}

type UpdateUserRequestDTO struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Input - Output Repository
type FindOneUserInput struct {
	Username string
	Password string
}
type FindUsersInput struct {
	Username string
	Age      string
}
type UpdateUserInput struct {
	Password string
	Email    string
}
type UpdateUserCondInput struct {
	Username string
	Email    string
	Age      int
}
