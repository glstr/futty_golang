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
