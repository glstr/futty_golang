package fos

import "testing"

func TestFinder(t *testing.T) {
	finder := NewFinder("./", MatchSuffix)
	result, err := finder.Find("go")
	if err != nil {
		t.Errorf("find failed, err_msg:%s", err.Error())
	}
	t.Logf("result:%v", result)
}
