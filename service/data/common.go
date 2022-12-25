package data

import (
	"errors"
	"reflect"
)

const (
	TagPrefix = "sql_tag"
)

var (
	ErrNotSupportType = errors.New("not support type")
)

// type Example struct {
//    Name string `sorm:"name"`
// }
func StructToMapVal(input interface{}) (map[string]interface{}, error) {
	inputType := reflect.TypeOf(input)
	inputVal := reflect.ValueOf(input)
	if inputType.Kind() == reflect.Ptr {
		inputType = inputType.Elem()
		inputVal = reflect.Indirect(inputVal)
	}

	if inputType.Kind() != reflect.Struct {
		return nil, ErrNotSupportType
	}

	tagMapName := make(map[string]string, inputType.NumField())
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		if field.IsExported() {
			tagKey := field.Tag.Get(TagPrefix)
			tagMapName[tagKey] = field.Name
		}
	}

	results := make(map[string]interface{}, len(tagMapName))
	for key, name := range tagMapName {
		val := inputVal.FieldByName(name)
		results[key] = val.Interface()
	}

	return results, nil
}

func StructToMapPoint(input interface{}) (map[string]interface{}, error) {
	inputType := reflect.TypeOf(input)
	if inputType.Kind() != reflect.Ptr {
		return nil, ErrNotSupportType
	}

	inputPointedType := inputType.Elem()
	if inputPointedType.Kind() != reflect.Struct {
		return nil, ErrNotSupportType
	}

	inputPointedVal := reflect.Indirect(reflect.ValueOf(input))

	tagMapName := make(map[string]string, inputPointedType.NumField())
	for i := 0; i < inputPointedType.NumField(); i++ {
		field := inputPointedType.Field(i)
		if field.IsExported() {
			tagKey := field.Tag.Get(TagPrefix)
			tagMapName[tagKey] = field.Name
		}
	}

	results := make(map[string]interface{}, len(tagMapName))
	for key, name := range tagMapName {
		val := inputPointedVal.FieldByName(name)
		results[key] = val.Addr().Interface()
	}

	return results, nil
}
