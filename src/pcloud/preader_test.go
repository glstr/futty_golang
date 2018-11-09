package pcloud

import (
	"fmt"
	"testing"
)

func TestPReaderLoad(t *testing.T) {
	//make data
	filePath := "/Users/pengbaojiang/pengbaojiang/code/gosrc/futty_golang/data/cloud.txt"
	tr := NewTxtReader(filePath)
	pcd := NewPointXYZCloud()
	tr.Load(pcd)
	pcd.Stat()
	fmt.Println(pcd.Points)
}
