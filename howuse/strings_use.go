package howuse

import (
	"log"
	"strings"
)

func SReplace(s string) string {
	return strings.Replace(s, "\n", "", -1)
}

func ShowTitleUse() string {
	s := "name_example"
	s = strings.Replace(s, "_", " ", -1)
	res := strings.Title(s)
	sArray := strings.Split(res, " ")
	res = strings.Join(sArray, "")
	log.Printf("%s", res)
	return res
}
