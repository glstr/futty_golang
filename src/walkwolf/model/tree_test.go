package model

import "testing"

func TestUrlTree(t *testing.T) {
	testUrl := "path/debug/hello"
	testUrl1 := "path/test/hello"
	testUrl2 := "path/debug/hello2"
	tree := NewUrlTree("root")
	err := tree.Insert(testUrl)
	err = tree.Insert(testUrl1)
	err = tree.Insert(testUrl2)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	t.Logf("str:%s", tree.String())
}
