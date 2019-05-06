package main

import (
	"flag"
	"log"
	"os"
)

const (
	pyHeader = "#!/usr/bin/python\n# coding:utf-8\n\n\n"
	mainFunc = "if __name__ == '__main__':\n    exit(1)"
)

var output = flag.String("out", "test.py", "help message for flagname")
var format = flag.String("format", "py", "file format")

func makePyFile(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = f.WriteString(pyHeader + mainFunc)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

func makeCppFile(className string) error {

}

func main() {
	flag.Parse()

	err := makePyFile(*output)
	if err != nil {
		log.Printf("err_msg:%s", err.Error())
	}
}
