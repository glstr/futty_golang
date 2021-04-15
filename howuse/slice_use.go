package howuse

import (
	"errors"
	"fmt"
)

//slice basic use
func MakeAndUse() {
	var e1 []int
	e1 = append(e1, 2)

	e2 := []int{1, 2, 3}

	fmt.Printf("e1:%v, e2:%v", e1, e2)
}

func ModifySlice() {
	e1 := []int{1, 2, 3, 4, 5}
	e2 := e1[:3]
	fmt.Printf("e1:%v, e2:%v", e1, e2)
	e2[0] = 6
	fmt.Printf("e1:%v, e2:%v", e1, e2)
	modify(e1)
	fmt.Printf("e1:%v, e2:%v", e1, e2)

	var e3 []int
	appendSlice(e3)
	fmt.Printf("e3:%v", e3)
}

func modify(i []int) {
	i[0] = 7
}

func appendSlice(i []int) {
	i = append(i, 4, 3)
}

func RemoveFront(l *([]int)) (int, error) {
	length := len(*l)
	var ret int
	if length == 0 {
		return ret, errors.New("input empty")
	}
	ret = (*l)[0]
	copy(*l, (*l)[1:])
	*l = (*l)[:length-1]
	return ret, nil
}
