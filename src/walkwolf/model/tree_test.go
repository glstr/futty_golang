package model

import "testing"

func TestUrlTree(t *testing.T) {
	testUrl := "path/debug/hello"
	testUrl1 := "path/test/hello"
	testUrl2 := "path/debug/hello2"
	testUrl3 := "path/test/hello"
	tree := NewUrlTree("root")
	err := tree.Insert(testUrl)
	err = tree.Insert(testUrl1)
	err = tree.Insert(testUrl2)
	err = tree.Insert(testUrl3)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	t.Logf("str:%s", tree.String())
}

func TestUrlTreeTraversal(t *testing.T) {
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

	res, err := tree.Traversal(0)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	t.Logf("res:%v", res)
}

func TestPreOrderTraversal(t *testing.T) {
	tree := makeTestTree()
	res, err := tree.root.preOrderTraversalEx()
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	t.Logf("res:%v", res)
}

func TestLevelOrderTravesal(t *testing.T) {
	tree := makeTestTree()
	res, err := tree.root.levelOrderTraversal()
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	t.Logf("res:%v", res)
}

func makeTestTree() *UrlTree {
	testUrl := "path/debug/hello"
	testUrl1 := "path/test/hello"
	testUrl2 := "path/debug/hello2"
	tree := NewUrlTree("root")
	tree.Insert(testUrl)
	tree.Insert(testUrl1)
	tree.Insert(testUrl2)
	return tree
}
