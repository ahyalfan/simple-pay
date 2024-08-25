package util

import (
	"github.com/go-playground/validator/v10"
)

func Vallidate[T any](data T) map[string]string {
	validate := validator.New()
	_ = validate.RegisterValidation("min1", Min1Float)
	err := validate.Struct(data)
	res := make(map[string]string)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			res[v.StructField()] = v.Error()
		}
	}
	return res
}

func Min1Float(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(float64)
	if ok {
		if value < 0.1 {
			return false
		}
		if value > 0.1 {
			return true
		}
	}
	return true
}
