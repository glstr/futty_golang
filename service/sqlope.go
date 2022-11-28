package service

func GetSqlOpeService() *SqlOpeService {
	return NewSqlOpeService()
}

type SqlOpeService struct{}

func NewSqlOpeService() *SqlOpeService {
	return &SqlOpeService{}
}

type SqlOpeReq struct{}

type SqlOpeRes struct {
	Result string
}

func (s *SqlOpeService) Ope() error {
	return nil
}
