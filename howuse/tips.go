package howuse

import "sync"

var (
	a     *int
	aOnce sync.Once
)

func getA() *int {
	aOnce.Do(func() {
		a := new(int)
		*a = 1
	})
	return a
}
