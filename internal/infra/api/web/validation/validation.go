package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
	"github.com/marcelockdata/go-leilao/configuration/rest_error"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		enTransl := ut.New(en, en)
		transl, _ = enTransl.GetTranslator("en")
		validator_en.RegisterDefaultTranslations(value, transl)
	}
}

func ValidateErr(validation_err error) *rest_error.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_error.NewNotFoundError("Invalid type error")
	} else if errors.As(validation_err, &jsonValidation) {
		errorCauses := []rest_error.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, rest_error.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			})
		}

		return rest_error.NewBadRequestError("Invalid field values", errorCauses...)
	} else {
		return rest_error.NewBadRequestError("Error trying to convert fields")
	}
}
