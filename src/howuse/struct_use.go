package howuse

import (
	"encoding/json"
	"fmt"
)

type BaseAnimal struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Panda struct {
	BaseAnimal
	From string
}

func MakePanda() {
	var panda Panda
	s, err := json.Marshal(panda)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(string(s))
}

//
type Programer struct {
	Name  string
	Level int
}

func (p *Programer) Coding() {
	fmt.Println("you can coding")
}

func (p *Programer) Debug() {
	fmt.Println("debug")
}

type SeniorProgramer struct {
	Programer
}

func (p *SeniorProgramer) Debug() {
	fmt.Println("find it")
}

func ProgramerWork() {
	var p SeniorProgramer
	fmt.Println(p)
	p.Debug()
}
