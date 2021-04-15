package model

import "testing"

func TestGetTargets(t *testing.T) {
	url := "http://www.xsnvshen.com/album/32456"
	target := &Target{
		RawUrl:  url,
		Element: "img",
		Attr:    "src",
	}

	parser := NewHTMLParser()
	t.Logf("url:%s", url)
	result, err := parser.GetTargetFromUrl(target)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	t.Logf("result:%v, len:%d", result, len(result))
}
