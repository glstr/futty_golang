package howuse

import (
	"fmt"
	"sync/atomic"
)

func Cas() {
	a := int32(3)
	atomic.CompareAndSwapInt32(&a, 3, 4)
	fmt.Println(a)

	b := atomic.LoadInt32(&a)
	fmt.Println(b)
}
