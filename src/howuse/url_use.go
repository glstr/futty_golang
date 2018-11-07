package howuse

import (
	"fmt"
	"net/url"
)

func QueryEscape() {
	e := "+ererer\n"
	temp := url.QueryEscape(e)
	fmt.Println(temp)
	res, _ := url.QueryUnescape(e)
	fmt.Printf("%s", res)
}
