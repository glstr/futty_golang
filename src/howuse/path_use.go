package howuse

import (
	"fmt"
	"path"
)

func GetDir(dirpath string) {
	fmt.Println(path.Dir(dirpath))
}
