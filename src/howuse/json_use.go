package howuse

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Decodejson() {
	example := "{\"hello\": 12343434343}"
	c := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(example))
	err := dec.Decode(&c)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c)
}
