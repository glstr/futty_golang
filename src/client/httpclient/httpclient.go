package httpclient

import (
	"io"
	"io/ioutil"
	"net/http"
)

var DefaultClient *HttpClient = NewHttpClient(true)

type HttpClient struct {
	client *http.Client
}

type Request struct {
	Method string
	Url    string
	Body   io.Reader
}

func NewHttpClient(isPool bool) *HttpClient {
	return &HttpClient{
		client: &http.Client{},
	}
}

func (c *HttpClient) Do(request *Request) ([]byte, error) {
	var body []byte
	req, err := http.NewRequest(request.Method, request.Url, request.Body)
	if err != nil {
		return body, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return body, err
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return body, err
	}
	return body, nil
}
