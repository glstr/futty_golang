package howuse

import "fmt"

func recoverBU() {
	fmt.Println("begin")
	for {
		defer func() {
			if e, ok := recover().(error); ok {
				fmt.Println(e)
			}
			go recoverBU()
		}()

		var tp map[string]int
		tp["hello"] = 1
	}
}

func RecoverBU() {
	go recoverBU()
}
