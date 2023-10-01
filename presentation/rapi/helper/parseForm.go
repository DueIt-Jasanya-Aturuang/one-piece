package helper

import (
	"net/http"
	"reflect"
)

func ParseForm(r *http.Request, data any) error {
	val := reflect.ValueOf(data).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("form")
		if tag != "" {
			if val.Field(i).Kind() == reflect.String {
				formField := r.FormValue(tag)
				val.Field(i).SetString(formField)
			}
		}
	}
	return nil
}
