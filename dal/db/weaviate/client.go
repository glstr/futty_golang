package weaviate

type Client interface {
}

func NewClient() Client {
	return &weaviateClient{}
}

type Config struct {
	URL    string
	Schema string
}

type weaviateClient struct {
}
