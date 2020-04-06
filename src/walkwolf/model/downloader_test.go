package model

import (
	"fmt"
	"testing"
)

func TestDownloader(t *testing.T) {
	downloader := NewDownloader()
	//for i := 0; i <= 40; i++ {
	url := fmt.Sprintf("https://img.xsnvshen.com/album/17204/32456/%.3d.jpg", 0)
	task := DownloadTask{
		URL: url,
		Dir: "/home/pi/pengbaojiang/code/gosrc/futty_golang/static/img",
	}
	err := downloader.RunTask(&task)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
	}
	//}
}
