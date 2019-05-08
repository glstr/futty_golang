package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	pyHeader = "#!/usr/bin/python\n# coding:utf-8\n\n\n"
	mainFunc = "if __name__ == '__main__':\n    exit(1)"
)

var output = flag.String("out", "test.py", "help message for flagname")
var format = flag.String("format", "py", "file format")

func makePyFile(fileName string) error {
	content := pyHeader + mainFunc
	return outputStringToFile(fileName, content)
}

func outputStringToFile(fileName, content string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}

// example_a
// file_name: example_a.h & example_a.cpp
// class_name: ExampleA
//
func makeContentInHeader(className string) string {
	content := "#pragma once\n\n"
	content += "using namespace snow {\n"
	content += "class %s {\n"
	content += "public:\n"
	content += "	%s (void);\n"
	content += "	virtual ~%s (void);\n"
	content += "};\n"
	content += "} //end namespace snow;"
	content = fmt.Sprintf(content, className, className, className)
	return content
}

func makeContentInCpp(headerFileName, className string) string {
	content := "#include \"%s\"\n\n"
	content += "namespace snow {\n"
	content += "%s::%s(void) {}\n\n"
	content += "%s::~%s(void) {}\n\n"
	content += "} //end namespace snow"
	content = fmt.Sprintf(content, headerFileName, className,
		className, className, className)
	return content
}

func makeClassName(fileName string) string {
	temp := strings.Replace(fileName, "_", " ", -1)
	temp = strings.Title(temp)
	eleStr := strings.Split(temp, " ")
	className := strings.Join(eleStr, "")
	return className
}

func makeCppFile(fileName string) error {
	className := makeClassName(fileName)
	fileHeaderName := fileName + ".h"
	err := outputStringToFile(fileHeaderName, makeContentInHeader(className))
	if err != nil {
		return err
	}

	fileCppName := fileName + ".cpp"
	err = outputStringToFile(fileCppName, makeContentInCpp(fileHeaderName, className))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()

	var err error
	if *format == "py" {
		err = makePyFile(*output)
	} else if *format == "cpp" {
		err = makeCppFile(*output)
	}

	if err != nil {
		log.Printf("err_msg:%s", err.Error())
	}
}
