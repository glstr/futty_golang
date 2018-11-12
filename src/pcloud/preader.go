package pcloud

import (
	"errors"
	"fmt"
	"os"
)

type PReader interface {
	Read(p PointCloud) (n int, err error)
	Load(p PointCloud) (n int, err error)
}

type TxtReader struct {
	FilePath string
}

func NewTxtReader(path string) *TxtReader {
	return &TxtReader{
		FilePath: path,
	}
}

func (r *TxtReader) Read(n int64, p PointCloud) (int64, error) {
	if pcd, ok := p.(*PointXYZCloud); !ok {
		return 0, errors.New(ErrWrongType)
	} else {
		f, err := os.Open(r.FilePath)
		if err != nil {
			return 0, err
		}
		defer f.Close()
		for i := int64(0); i < n; i++ {
			var p PointXYZ
			_, err := fmt.Fscanf(f, "%f %f %f\n", &p.X, &p.Y, &p.Z)
			if err != nil {
				return i, errors.New(ErrNomoreData)
			}
			pcd.Points = append(pcd.Points, p)
		}
		return n, nil

	}
}

func (r *TxtReader) Load(p PointCloud) error {
	if pcd, ok := p.(*PointXYZCloud); !ok {
		return errors.New(ErrWrongType)
	} else {
		f, err := os.Open(r.FilePath)
		if err != nil {
			return err
		}
		defer f.Close()
		for {
			var p PointXYZ
			_, err := fmt.Fscanf(f, "%f %f %f\n", &p.X, &p.Y, &p.Z)
			if err != nil {
				return err
			}
			pcd.Points = append(pcd.Points, p)
		}
		return nil
	}
}

type BlsReader struct {
	FilePath string
}

func NewBlsReader(path string) *BlsReader {
	return &BlsReader{
		FilePath: path,
	}
}

func (r *BlsReader) Read(n int64, p PointCloud) (int64, error) {
	return 0, nil
}

func (r *BlsReader) Load(p PointCloud) error {
	return nil
}

//las reader:
//lasfile: one file format for point cloud
//support las 1.2 or 1.3
type LasReader struct {
	FilePath string
}

func NewLasReader(path string) *LasReader {
	return &LasReader{
		FilePath: path,
	}
}

func (r *LasReader) Read(n int64, p PointCloud) (int64, error) {
	return 0, nil
}

func (r *LasReader) Load(p PointCloud) error {
	return nil
}
