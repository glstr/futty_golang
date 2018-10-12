package howuse

import "fmt"

// func wrap
type FuncA func(a, b int) int
type FuncB func(a, b float32) float32

func GeneFunc(fa FuncA) FuncB {
	return func(a, b float32) float32 {
		return float32(fa(int(a), int(b)))
	}
}

func GeneFuncEx(fa interface{}) FuncB {
	//
	return func(a, b float32) float32 {
		fmt.Println("see reflect_use.go")
		return 0.0
	}
}

// closure
func ClosureMU() {
	var a, b int32
	a, b = 32, 64
	fmt.Printf("outspace:%v, %v\n", &a, &b)
	c := make(chan struct{})
	go func(a, b int32) {
		fmt.Printf("outspace:%v, %v\n", &a, &b)
		c <- struct{}{}
	}(a, b)
	<-c
}
