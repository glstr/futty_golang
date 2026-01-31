package nettool

import "testing"

func TestIP2Region(t *testing.T) {
	ip := "110.242.70.57"
	info, err := IP2Region(ip)
	if err != nil {
		t.Errorf("trans failed:%s", err.Error())
		return
	}

	t.Logf("info:%v", info)
}
