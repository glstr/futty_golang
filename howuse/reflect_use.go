package howuse

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

//basic usage
type Animal interface {
	Roar()
}

type Cat struct {
	Name string
	Age  int32
}

func (c *Cat) Roar() {
	fmt.Println("It is a cat")
}

func RfMu() {
	//reflect.Type
	var animal Animal
	cat := &Cat{}
	animal = cat
	t := reflect.TypeOf(animal).Elem()
	fmt.Println(t.Name())
	fmt.Println(t.Kind())

	tt := reflect.TypeOf(t).Elem()
	fmt.Println(tt.Name())
	fmt.Println(tt.Kind())

	//reflect.Value
	v := reflect.ValueOf(animal).Elem()
	v.Field(0).SetString("hello")
	v.Field(1).SetInt(32)
	fmt.Println(animal)
}

func RfEleMU() {
	m := map[string]string{"hello": "world"}

	t := reflect.TypeOf(m).Elem()
	v := reflect.ValueOf(m)

	fmt.Println(t.Kind())
	fmt.Println(v.Kind())

	example := func(tm map[string]string) {
		tmv := reflect.ValueOf(tm)
		tmt := reflect.TypeOf(tm)

		kt := tmt.Key()
		vt := tmt.Elem()

		ktv := reflect.New(kt).Elem()
		vtv := reflect.New(vt).Elem()

		ktv.SetString("snow")
		vtv.SetString("window")

		tmv.SetMapIndex(ktv, vtv)
	}
	example(m)
	fmt.Println(m)
}

//map to struct && struct to map

type Singer struct {
	Name string
	Age  int
}

func ReflectUse() {
	singer := Singer{
		Name: "Jim",
		Age:  1,
	}

	t := reflect.TypeOf(singer)
	fmt.Println(t.Kind())
	v := reflect.ValueOf(singer)
	num := t.NumField()
	res := make(map[string]interface{})
	for i := 0; i < num; i++ {
		field := t.Field(i)
		name := field.Name
		val := v.FieldByName(name)
		res[name] = val.Interface()
	}
	fmt.Println(res)
	ParseMap(res)
}

func structToMap(i interface{}) map[string]interface{} {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if reflect.TypeOf(i).Kind() != reflect.Struct {
		fmt.Println("param error")
		return nil
	}
	num := t.NumField()
	res := make(map[string]interface{}, num)
	for i := 0; i < num; i++ {
		name := t.Field(i).Name
		val := v.FieldByName(name)
		res[name] = val.Interface()
	}
	return res
}

func StructToMap() {
	singer := Singer{
		Name: "Lucy",
		Age:  4,
	}
	res := structToMap(singer)
	fmt.Println(res)
}

// reflect usage of function
// modify function
// target funcA { funcB }
// funcA := wrap(funcB)
type OldFunc func(a, b int32) int32

func MakeFunc() {
	examfunc := func(a, b int32) int32 {
		return int32(a - b)
	}

	oldFunc, err := makeFunc(examfunc)
	if err != nil {
		fmt.Println(err)
	}
	ret := oldFunc(2, 3)
	fmt.Println(ret)
}

func makeFunc(action interface{}) (OldFunc, error) {
	t := reflect.TypeOf(action)
	if t.Kind() != reflect.Func {
		return nil, errors.New(PARAM_ERROR)
	}

	wrap := func(args []reflect.Value) []reflect.Value {
		fmt.Println(args)
		a, b, err := prepare(args)
		fmt.Printf("a:%d, b:%d, err:%v", a, b, err)

		oa := reflect.New(t.In(0)).Elem()
		ob := reflect.New(t.In(1)).Elem()

		oa.SetInt(int64(a))
		ob.SetInt(int64(b))
		params := []reflect.Value{
			oa,
			ob,
		}
		ret := reflect.ValueOf(action).Call(params)
		var res []reflect.Value
		res = append(res, ret[0])
		return res
	}

	//make new func
	var f OldFunc
	ft := reflect.ValueOf(&f).Elem()
	fv := reflect.MakeFunc(ft.Type(), wrap)
	ft.Set(fv)
	return f, nil
}

func prepare(args []reflect.Value) (int32, int32, error) {
	var a int32
	var b int32
	if len(args) != 2 {
		return a, b, errors.New(PARAM_ERROR)
	}

	if !args[0].CanInterface() || !args[1].CanInterface() {
		return a, b, errors.New(PARAM_ERROR)
	}

	i1 := args[0].Interface()
	i2 := args[1].Interface()

	var ok bool
	if a, ok = i1.(int32); !ok {
		return a, b, errors.New(PARAM_ERROR)
	}

	if b, ok = i2.(int32); !ok {
		return a, b, errors.New(PARAM_ERROR)
	}

	return a, b, nil
}

// reflect like func template
func AddInt(a, b int) int {
	return a + b
}

func AddFloat(a, b float32) float32 {
	return a + b
}

type AddIntType func(int, int) int
type AddFloatType func(float32, float32) float32

func implement(in []reflect.Value) []reflect.Value {
	addInt := func(a, b int) int {
		return a + b
	}

	addFloat := func(a, b float32) float32 {
		return a + b
	}

	if len(in) != 2 {
		return []reflect.Value{}
	}

	if in[0].Type().Kind() == reflect.Float32 {
		return reflect.ValueOf(addFloat).Call(in)
	} else if in[0].Type().Kind() == reflect.Int {
		return reflect.ValueOf(addInt).Call(in)
	} else {
		return []reflect.Value{}
	}
}

func getNewFunc(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), implement)
	fn.Set(v)
}

func MafunMU() {
	var f AddIntType
	getNewFunc(&f)
	ret := f(2, 3)
	fmt.Println(ret)
}

//usage of SelectCase
func ShowSelectCase() {
	//env
	c1 := make(chan int)
	c2 := make(chan int)
	f := func(c chan int) {
		for i := 0; i < 10; i++ {
			c <- rand.Intn(100)
			time.Sleep(1 * time.Second)
		}
		close(c)
	}
	go f(c1)
	go f(c2)
	case1 := reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(c1),
	}
	case2 := reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(c2),
	}
	var cases []reflect.SelectCase
	cases = append(cases, case1)
	cases = append(cases, case2)

	for {
		chosen, recv, recvOk := reflect.Select(cases)
		fmt.Printf("chosen:%d\n", chosen)
		if !recvOk {
			fmt.Println("recv error")
			break
		}
		if !recv.CanInterface() {
			fmt.Println("value error")
			break
		}
		v, ok := recv.Interface().(int)
		if !ok {
			fmt.Println("value trans error")
			break
		}
		fmt.Printf("num:%d\n", v)
	}
}
