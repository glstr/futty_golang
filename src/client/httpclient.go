/*
Package model provides examples of useful tool, such as ratelimiter, task_group and so on.
*/
package model

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type config struct {
	Method string
	Url    string
	Body   io.Reader
}

var reqCount int64

func makeClient() *http.Client {
	var defaultTransport RoundTripper = &Transport{
		Proxy: ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{
		Transport: defaultTransport,
	}
	return client
}

func makeRequest(c *config) {
	request := http.NewRequest(c.Method, c.Url, c.Body)
	return request
}

func doRequest(client *http.Client, c *config) {
	for {
		req = makeRequest(c)
		res, err := client.Do(req)
		if err != nil {
			log.Printf("[errmsg:%s]", err.Error())
		}
		defer res.body.Close()
		content, err := ioutil.ReadAll(res.body)
		if err != nil {
			log.Printf("[content:%s]", string(content))
		}
		time.Sleep(1 * time.Millisecond)
		atomic.AddInt64(&reqCount, 1)
	}
}

func start() {
	client := makeClient()
	c := &config{
		Method: "Post",
		Url:    "test",
		Body:   nil,
	}
	max := 10
	for i := 0; i < 10; i++ {
		go doRequest(client, c)
	}
	go startStat()
}

func startStat() {
	var lastCount int64
	for {
		curr := atomic.LoadInt64(&reqCount)
		qps := curr - lastCount
		time.Sleep(1 * time.Second)
		log.Printf("[qps:%d]", qps)
	}
}
