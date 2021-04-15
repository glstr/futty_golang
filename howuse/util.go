package howuse

import "fmt"

const (
	PARAM_ERROR = "param error"
)

func CheckError(err error) bool {
	if err != nil {
		fmt.Printf("errMsg:%v", err)
		return false
	}
	return true
}
