package howuse

import "fmt"

//basic usage
func MapMakeUse() {
	example := map[string]int{"hello": 1, "world": 2}

	if v, ok := example["hi"]; ok {
		fmt.Printf("you get it:%d\n", v)
	} else {
		fmt.Printf("nothing")
	}
}

//map[string]interface
func ParseMap(e map[string]interface{}) {
	for k, v := range e {
		switch vv := v.(type) {
		case string:
			fmt.Printf("key:%s, string:%s\n", k, vv)
		case int:
			fmt.Printf("key:%s, string:%d\n", k, vv)
		case float64:
			fmt.Printf("key:%s, float:%f\n", k, vv)
		default:
			fmt.Printf("unknown")
		}
	}
}
