package howuse

import (
	"fmt"
	"reflect"
)

//map to struct && struct to map

type Singer struct {
	Name string
	Age  int
}

func ReflectUse() {
	singer := Singer{
		Name: "Jim",
		Age:  1,
	}

	t := reflect.TypeOf(singer)
	fmt.Println(t.Kind())
	v := reflect.ValueOf(singer)
	num := t.NumField()
	res := make(map[string]interface{})
	for i := 0; i < num; i++ {
		field := t.Field(i)
		name := field.Name
		val := v.FieldByName(name)
		res[name] = val.Interface()
	}
	fmt.Println(res)
	ParseMap(res)
}

func structToMap(i interface{}) map[string]interface{} {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if reflect.TypeOf(i).Kind() != reflect.Struct {
		fmt.Println("param error")
		return nil
	}
	num := t.NumField()
	res := make(map[string]interface{}, num)
	for i := 0; i < num; i++ {
		name := t.Field(i).Name
		val := v.FieldByName(name)
		res[name] = val.Interface()
	}
	return res
}

func StructToMap() {
	singer := Singer{
		Name: "Lucy",
		Age:  4,
	}
	res := structToMap(singer)
	fmt.Println(res)
}
