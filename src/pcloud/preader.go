package pcloud

import (
	"errors"
	"fmt"
	"os"
)

type PointCloud interface {
	Stat()
}

type PointXYZ struct {
	X float64
	Y float64
	Z float64
}

func (p *PointXYZ) Stat() {
}

type PReader interface {
	Read(p PointCloud) (n int, err error)
}

type TxtReader struct {
	FilePath string
}

func NewTxtReader(path string) *TxtReader {
	return &TxtReader{
		FilePath: path,
	}
}

func (r *TxtReader) Read(p PointCloud) (n int, err error) {

	if pcd, ok := p.(*PointXYZ); !ok {
		return 0, errors.New("wrong type")
	} else {
		f, err := os.Open(r.FilePath)
		if err != nil {
			return 0, err
		}
		fmt.Fscanf(f, "%f %f %f\n", &pcd.X, &pcd.Y, &pcd.Z)
		fmt.Println(pcd)
	}
	return 0, nil
}
