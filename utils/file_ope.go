package utils

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func GetTargetFilesFromDir(dir string, suffix string) []string {
	var result []string
	f := func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(path, suffix) {
			newPath := strings.TrimPrefix(path, dir)
			result = append(result, newPath)
		}
		return nil
	}

	filepath.WalkDir(dir, f)
	return result
}
