package models

import "reflect"

var (
	Helpers *_Helpers
)

type _Helpers struct{}

func (_ *_Helpers) ToMapFromString(from reflect.Kind, to reflect.Kind, data interface{}) (interface{}, error) {
	if (from == reflect.String) && (to == reflect.Map) {
		return map[string]interface{}{}, nil
	}

	return data, nil
}
