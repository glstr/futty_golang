package resource

import "time"

type Video struct {
	Name     string
	Author   string
	Duration int64
	Tags     []string

	// file real location
	Path       string
	CreateTime time.Time
	UpdateTime time.Time
}

func NewVideo() *Video {
	return &Video{}
}

type VideoManager struct{}
