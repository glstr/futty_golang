package rpctest

import (
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	res, err := http.Get("www.baidu.com")
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}

	t.Logf("res:%v", res)
}
