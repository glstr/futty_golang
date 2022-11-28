package spider

type Spider struct {
	limit int64
}

func NewSpider() *Spider {
	return &Spider{
		limit: 1,
	}
}

func (s *Spider) Crawl(oriAddr string) error {
	return nil
}
