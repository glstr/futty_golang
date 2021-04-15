package pcloud

import "fmt"

//error code & error type
var (
	ErrWrongType  = "wrong type"
	ErrNomoreData = "no more data"
)

var (
	defaultSize int64 = 0
)

type PointXYZ struct {
	X float64
	Y float64
	Z float64
}

type PointCloud interface {
	Stat()
}

type PointXYZCloud struct {
	Points []PointXYZ
}

func NewPointXYZCloud() *PointXYZCloud {
	return &PointXYZCloud{
		Points: make([]PointXYZ, defaultSize),
	}
}

func (p *PointXYZCloud) Stat() {
	size := len(p.Points)
	fmt.Printf("points size:%d", size)
}
