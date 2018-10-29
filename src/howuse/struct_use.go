package howuse

import (
	"encoding/json"
	"fmt"
)

type BaseAnimal struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Panda struct {
	BaseAnimal
	From string
}

func MakePanda() {
	var panda Panda
	s, err := json.Marshal(panda)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(string(s))
}
