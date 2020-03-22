package action

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

const (
	FormatPy  = "py"
	FormatCpp = "cpp"
)

const (
	pyHeader = "#!/usr/bin/python\n# coding:utf-8\n\n\n"
	mainFunc = "if __name__ == '__main__':\n    exit(1)"
)

var (
	ErrFormatUnsupport = errors.New("format not support")
)

func CodeGene(c *cli.Context) error {
	var (
		keyFormat   = "format"
		keyFileName = "filename"
	)

	format := c.String(keyFormat)
	fileName := c.String(keyFileName)
	log.Printf("format:%s, file_name:%s", format, fileName)
	m := GetMaker(format)
	if m == nil {
		return ErrFormatUnsupport
	}

	return m.Make(fileName)
}

func GetMaker(format string) FileMaker {
	switch format {
	case FormatPy:
		return &PythonMaker{}
	case FormatCpp:
		return &CppMaker{}
	default:
		return nil
	}
}

type FileMaker interface {
	Make(filePath string) error
}

type PythonMaker struct{}

func (pm *PythonMaker) Make(fileName string) error {
	content := pm.makeDefaultContent()
	filePath := pm.makeFilePath(fileName)
	return outputStringToFile(filePath, content)
}

func (pm *PythonMaker) makeFilePath(fileName string) string {
	return fileName + ".py"
}

func (pm *PythonMaker) makeDefaultContent() string {
	return pyHeader + mainFunc
}

func outputStringToFile(filePath, content string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

type CppMaker struct{}

func (cm *CppMaker) Make(fileName string) error {
	err := cm.makeHeaderFile(fileName)
	if err != nil {
		return err
	}
	return cm.makeCppFile(fileName)
}

func (cm *CppMaker) makeHeaderFile(fileName string) error {
	filePath := cm.makeHeaderFilePath(fileName)
	className := cm.makeClassName(fileName)
	content := cm.makeContentInHeader(className)
	return outputStringToFile(filePath, content)
}

func (cm *CppMaker) makeCppFile(fileName string) error {
	filePath := cm.makeCppFilePath(fileName)
	className := cm.makeClassName(fileName)
	content := cm.makeContentInHeader(className)
	return outputStringToFile(filePath, content)
}

func (cm *CppMaker) makeContentInHeader(className string) string {
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

func (cm *CppMaker) makeContentInCpp(headerFileName, className string) string {
	content := "#include \"%s\"\n\n"
	content += "namespace snow {\n"
	content += "%s::%s(void) {}\n\n"
	content += "%s::~%s(void) {}\n\n"
	content += "} //end namespace snow"
	content = fmt.Sprintf(content, headerFileName, className,
		className, className, className)
	return content
}

func (cm *CppMaker) makeClassName(fileName string) string {
	temp := strings.Replace(fileName, "_", " ", -1)
	temp = strings.Title(temp)
	eleStr := strings.Split(temp, " ")
	className := strings.Join(eleStr, "")
	return className
}

func (cm *CppMaker) makeHeaderFilePath(fileName string) string {
	return fileName + ".h"
}

func (cm *CppMaker) makeCppFilePath(fileName string) string {
	return fileName + ".cpp"
}
