package httpclient

import (
	"bytes"
	"testing"
)

func TestHttpClient(t *testing.T) {
	method := "GET"
	url := "http://www.baidu.com"
	req := &Request{
		Method: method,
		Url:    url,
		Body:   bytes.NewBuffer([]byte("")),
	}

	body, err := DefaultClient.Do(req)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}

	t.Logf("res:%s", string(body))
}
