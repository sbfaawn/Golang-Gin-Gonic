package validator

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterValidation("notblank", NotBlank)
	Validate.RegisterValidation("password", ValidPassword)
}

func NotBlank(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(field.String())) > 0
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

func ValidPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	var (
		hasNumber      = false
		hasLetter      = false
		suitableLength = false
		hasUpper       = false
		hasLower       = false
	)

	if len(password) > 6 {
		suitableLength = true
	}

	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsLetter(c) || c == ' ':
			hasLetter = true
			if unicode.IsUpper(c) {
				hasUpper = true
			}

			if unicode.IsLower(c) {
				hasLower = true
			}
		default:
			return false
		}
	}

	return hasNumber && hasLetter && suitableLength && hasUpper && hasLower
}
