package model

import "testing"

func TestGetTargets(t *testing.T) {
	url := "http://www.xsnvshen.com/album/32456"
	target := "img"
	//target := "a"
	parser := &HTMLParser{}
	t.Logf("url:%s", url)
	result, err := parser.GetTargets(url, target)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	t.Logf("result:%v, len:%d", result, len(result))
}
