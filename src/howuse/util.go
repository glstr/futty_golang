package howuse

import "fmt"

func CheckError(err error) bool {
	if err != nil {
		fmt.Printf("errMsg:%v", err)
		return false
	}
	return true
}
