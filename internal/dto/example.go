package dto

// Request - Response UseCase
type ValidateExampleRequestDTO struct {
	ValidateCustom string `json:"validateCheck" binding:"required,customvalidate"`
	ValidateStruct int    `json:"validateStruct" binding:"required" validate:"gte=10,lte=20"`
}
