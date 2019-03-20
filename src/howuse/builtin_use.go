package howuse

import "fmt"

// elegant constant
const (
	MDog  = iota //0
	MCat         //1
	MBird        //2
)

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

const MaxUint32 = 1<<32 - 1
const MaxUint = ^uint(0)

// constant
// '' means rune -> uint32
var f = 'a' * 1.5

func CMU() {
	fmt.Println(f)
}

//type
type Header map[string][]string

func ShowTypeUse() {
	header := make(Header)
	header["First"] = []string{"hello", "world"}
	header["Second"] = []string{"Jim", "Lily"}
	fmt.Printf("header:%v", header)
}

type HandleFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

func ShowTypeUseFunc() {
	var f HandleFunc
	f = add
	res := f(1, 2)
	fmt.Printf("%d", res)
}
