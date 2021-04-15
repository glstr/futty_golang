package utils

import (
	"errors"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"sync"
	"time"
)

func StructToMapValue(input interface{}) (map[string]interface{}, error) {
	inputType := reflect.TypeOf(input)
	if inputType.Kind() != reflect.Struct {
		return nil, errors.New("input type error")
	}
	inputValue := reflect.ValueOf(input)

	numField := inputType.NumField()
	res := make(map[string]interface{})
	for i := 0; i < numField; i++ {
		if inputValue.Field(i).CanInterface() {
			key := inputType.Field(i).Tag.Get("json")
			value := inputValue.Field(i).Interface()
			res[key] = value
		}
	}
	return res, nil
}

func StructToMapAddr(input interface{}) (map[string]interface{}, error) {
	inputType := reflect.TypeOf(input)
	if inputType.Kind() != reflect.Ptr {
		return nil, errors.New("input type error")
	}
	eleType := inputType.Elem()
	if eleType.Kind() != reflect.Struct {
		return nil, errors.New("input type error")
	}

	inputValue := reflect.ValueOf(input).Elem()
	numField := eleType.NumField()
	res := make(map[string]interface{})
	for i := 0; i < numField; i++ {
		if inputValue.Field(i).CanAddr() {
			key := eleType.Field(i).Tag.Get("json")
			addr := inputValue.Field(i).Addr().Interface()
			res[key] = addr
		}
	}
	return res, nil
}

func GetIntFromInteface(input interface{}) (int, error) {
	switch t := input.(type) {
	case string:
		val, err := strconv.Atoi(t)
		if err != nil {
			return 0, err
		}
		return val, nil
	case int:
		return t, nil
	default:
		return 0, errors.New("unkown type")
	}
}

type Displayer struct {
	mutex sync.Mutex
	index int32
}

func NewDisplayer() *Displayer {
	return &Displayer{
		index: 0,
	}
}

func (d *Displayer) display() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	log.Printf("%d", d.index)
	d.index++
	return
}

func (d *Displayer) End() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	log.Printf("done:%d", d.index)
	d.index++
	return
}

func DisplayerFunc(done chan struct{}, d *Displayer) {
	for {
		select {
		case <-done:
			d.End()
			return
		case <-time.After(1 * time.Second):
			d.display()
		}
	}
}

type ShowData struct {
	Value int64  `json:"value"`
	Date  string `json:"date"`
}

func GenerateData() ShowData {
	date := time.Now().Format("15:04:05")
	return ShowData{
		Value: rand.Int63n(120),
		Date:  date,
	}
}
