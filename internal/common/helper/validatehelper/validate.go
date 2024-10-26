package validatehelper

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type (
	ValidateHelper interface {
		ValidateStruct(data any) error
	}

	validateHelper struct {
		validate *validator.Validate
	}
)

var CustomValidate validator.Func = func(fl validator.FieldLevel) bool {
	data := fl.Field().String()
	if data == "customValidate" {
		return true
	}
	return false
}

func NewValidate() ValidateHelper {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("customvalidate", CustomValidate)
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	return &validateHelper{
		validate: validate,
	}
}

func (h *validateHelper) ValidateStruct(data any) error {
	return h.validate.Struct(data)
}
