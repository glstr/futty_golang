package model

import "testing"

func TestUrlHelper(t *testing.T) {
	rawPath := "http://127.0.0.1:6060/pkg/strings/#Split"
	uh, err := ParseUrl(rawPath)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
	}
	t.Logf("path_section:%v", uh.GetPathSection())
	t.Logf("pure_path:%s", uh.GetPurePath())
}
