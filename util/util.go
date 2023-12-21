package util

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type Token struct{}

// ExtractBearerToken Remove "Bearer " from "Authorization" token string
func (tk Token) ExtractBearerToken(t string) (string, error) {
	f := strings.Split(t, " ")
	if len(f) != 2 || f[0] != "Bearer" {
		return "", nil
	}
	return f[1], nil
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(str interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(str)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
