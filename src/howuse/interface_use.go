package howuse

import (
	"fmt"
	"strconv"
)

type Iface interface {
	Hello()
}

type Func func(a, b int) string

func (f *Func) Hello() {
	res := (*f)(2, 3)
	fmt.Println(res)
}

type Dog struct {
	Name string
}

func (d *Dog) Hello() {
	fmt.Println("dog")
	d.Name = "Window"
}

func IfaceUse() {
	example := func(a, b int) string {
		c := a + b
		s := strconv.Itoa(c)
		return s
	}
	var funct Func
	funct = example
	funct.Hello()

	var dog Dog
	dog.Hello()
	fmt.Println(dog)

	var i Iface
	var j Iface
	i = &dog
	j = &funct
	i.Hello()
	j.Hello()
}
