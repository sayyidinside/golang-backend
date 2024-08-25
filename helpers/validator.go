package helpers

import (
	"errors"
	"golang/backend/dtos"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func validateStruct(param any) []*ErrorResponse {
	var errs []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(param)
	if err != nil {
		v := reflect.ValueOf(param)
		t := v.Type()

		for _, err := range err.(validator.ValidationErrors) {
			field, _ := t.FieldByName(err.StructField())
			formName := field.Tag.Get("form")
			if formName == "" {
				formName = err.StructField()
			}

			element := ErrorResponse{
				Field: formName,
				Tag:   err.Tag(),
				Value: err.Param(),
			}
			errs = append(errs, &element)
		}
	}
	return errs
}

func ValidateInput(input interface{}) (*dtos.Response, error) {
	if errs := validateStruct(input); len(errs) > 0 {
		return &dtos.Response{
			Success: false,
			Message: "Validation failed",
			Error:   errs,
		}, errors.New("validation errors occurred")
	}

	return nil, nil
}
