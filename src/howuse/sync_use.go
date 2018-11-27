package howuse

import (
	"fmt"
	"sync"
)

func ShowWGUse() {
	var wg sync.WaitGroup
	n := 10000
	res := make([]int64, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(input int) {
			defer wg.Done()
			res[input] = int64(input) * int64(input)
		}(i)
	}
	wg.Wait()
	var sum int64 = 0
	for _, v := range res {
		sum = sum + v
	}
	fmt.Printf("sum:%d", sum)
}
