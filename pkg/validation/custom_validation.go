package validation

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func NotEmptySlice(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() == reflect.Slice {
		return field.Len() > 0
	}

	return false
}
