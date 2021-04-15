package howuse

import (
	"fmt"
	"os"
)

func CFile(filePath string) {
	_, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
	}
}

func MakeDir(dirPath string) {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
