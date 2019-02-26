package utils

import "errors"

func MapToStruct(in map[string][]string, out interface{}) error {
	if len(in) <= 0 {
		return errors.New("in is empty")
	}
	//parse out, get json tag and field type
	return nil
}
