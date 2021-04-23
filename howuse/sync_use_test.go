/*
Package howuse show basic usage of go and some go package. In this package you can get some basic examples and patterns for
some go func and Package
*/

package howuse

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
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
}

//ShowCondUse show basic usage of cond
func ShowCondUse() {
	c := sync.NewCond(&sync.Mutex{})
	c.L.Lock()
	for true {
		c.Wait()
	}
	c.L.Unlock()
}

//ShowPoolUse show basic usage of pool
func ShowPoolUse() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new instance")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}

//Condition
