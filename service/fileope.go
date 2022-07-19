package service

type FileService interface {
	Upload(input string) (string, error)
	Download(input string) error
}

type DefaultFileService struct{}
