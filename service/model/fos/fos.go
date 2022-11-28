package fos

import (
	"errors"
	"io/fs"
	"path/filepath"
	"strings"
)

var (
	ErrNotFound = errors.New("not found")
)

// Watcher can find all file in pointed directory and sort of them;
type Watcher struct {
	FilesKinds []string
	KindsCount map[string]int64
}

type MatchFunc func(path string, pattern string) bool

func MatchSuffix(path string, pattern string) bool {
	return strings.HasSuffix(path, pattern)
}

func MatchPrefix(path string, pattern string) bool {
	return strings.HasPrefix(path, pattern)
}

type Finder struct {
	Dir       string
	matchFunc MatchFunc
}

func NewFinder(dir string, fn MatchFunc) *Finder {
	return &Finder{
		Dir:       dir,
		matchFunc: fn,
	}
}

func (f *Finder) Find(pattern string) ([]string, error) {
	var result []string
	fn := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if f.matchFunc != nil {
			if f.matchFunc(path, pattern) {
				result = append(result, path)
			}
		}

		return nil
	}

	err := filepath.WalkDir(f.Dir, fn)
	if err != nil {
		return nil, err
	}

	return result, nil
}
