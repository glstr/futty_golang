package service

import (
	"sync"

	"github.com/glstr/futty_golang/model"
	"github.com/glstr/futty_golang/utils"
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
	videos := utils.GetTargetFilesFromDir("/media/pi/glstr/download/download", "mp4")
	if videoId > int64(len(videos)) {
		return nil, model.ErrNotFound
	}

	return &VideoInfo{videoId, videos[videoId], ""}, nil
}

func (m *DefaultVideoService) GetVideoList() ([]*VideoInfo, error) {
	videos := utils.GetTargetFilesFromDir("/media/pi/glstr/download/download", "mp4")
	var res []*VideoInfo
	for id, path := range videos {
		info := &VideoInfo{
			Id:   int64(id),
			Path: path,
		}
		res = append(res, info)
	}
	return res, nil
}
