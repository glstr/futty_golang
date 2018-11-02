package howuse

import "strings"

func SReplace(s string) string {
	return strings.Replace(s, "\n", "", -1)
}
