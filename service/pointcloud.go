package service

type PCloudService interface {
	Process(input, output string) error
}

func GetPCloudService(serType int) PCloudService {
	return &SnowPCloudService{}
}

type SnowPCloudService struct{}

func (s *SnowPCloudService) Process(input string, output string) error {
	return nil
}
