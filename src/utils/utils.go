package utils

import (
	"errors"
	"reflect"
)

func StructToMapValue(input interface{}) (map[string]interface{}, error) {
	inputType := reflect.TypeOf(input)
	if inputType.Kind() != reflect.Struct {
		return nil, errors.New("input type error")
	}
	inputValue := reflect.ValueOf(input)

	numField := inputType.NumField()
	res := make(map[string]interface{})
	for i := 0; i < numField; i++ {
		if inputValue.Field(i).CanInterface() {
			key := inputType.Field(i).Tag.Get("json")
			value := inputValue.Field(i).Interface()
			res[key] = value
		}
	}
	return res, nil
}

func StructToMapAddr(input interface{}) (map[string]interface{}, error) {
	inputType := reflect.TypeOf(input)
	if inputType.Kind() != reflect.Ptr {
		return nil, errors.New("input type error")
	}
	eleType := inputType.Elem()
	if eleType.Kind() != reflect.Struct {
		return nil, errors.New("input type error")
	}

	inputValue := reflect.ValueOf(input).Elem()
	numField := eleType.NumField()
	res := make(map[string]interface{})
	for i := 0; i < numField; i++ {
		if inputValue.Field(i).CanAddr() {
			key := eleType.Field(i).Tag.Get("json")
			addr := inputValue.Field(i).Addr().Interface()
			res[key] = addr
		}
	}
	return res, nil
}
