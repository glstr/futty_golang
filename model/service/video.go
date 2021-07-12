package service

import (
	"sync"

	"github.com/glstr/futty_golang/model"
)

var (
	videos = map[int64]string{
		1: "/SNIS-104.avi",
		2: "/SSNI-658.mp4",
		3: "/SSNI-378.mp4",
		4: "/ssni-567/ssni-567.mp4",
		5: "/tk/0675.mp4",
		6: "/999.mp4",
	}
)

type VideoInfo struct {
	Id   int64  `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

type VideoService interface {
	GetVideo(videoId int64) (*VideoInfo, error)
	GetVideoList() ([]*VideoInfo, error)
}

var (
	videoServiceOnce sync.Once
	videoService     VideoService
)

func GetVideoService() VideoService {
	videoServiceOnce.Do(func() {
		videoService = &DefaultVideoService{}
	})
	return videoService
}

type DefaultVideoService struct{}

func (m *DefaultVideoService) GetVideo(videoId int64) (*VideoInfo, error) {
	path, ok := videos[videoId]
	if !ok {
		return nil, model.ErrNotFound
	}
	return &VideoInfo{videoId, path, ""}, nil
}

func (m *DefaultVideoService) GetVideoList() ([]*VideoInfo, error) {
	var res []*VideoInfo
	for id, path := range videos {
		info := &VideoInfo{
			Id:   id,
			Path: path,
		}
		res = append(res, info)
	}
	return res, nil
}
