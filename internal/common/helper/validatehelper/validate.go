package validatehelper

import "github.com/go-playground/validator/v10"

type (
	ValidateHelper interface {
		ValidateStruct(data any) error
	}

	validateHelper struct {
		validate *validator.Validate
	}
)

func NewValidate() ValidateHelper {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return &validateHelper{
		validate: validate,
	}
}

func (h *validateHelper) ValidateStruct(data any) error {
	return h.validate.Struct(data)
}
