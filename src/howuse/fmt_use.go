package howuse

import (
	"fmt"
	"os"
)

func WriteData() {
	fp, err := os.Create("text.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprintf(fp, "%d %d %d", 1, 2, 3)
	fp.Close()

	fp, err = os.Open("text.txt")
	var a, b, c int64
	fmt.Fscanf(fp, "%d %d %d", &a, &b, &c)
	fmt.Println(a, b, c)
}
