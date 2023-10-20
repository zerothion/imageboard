package utils

import "reflect"

func IsAnyFieldNil(obj interface{}) bool {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsNil() {
			return true
		}
	}

	return false
}
