package action

import (
	"client/httpclient"
	"log"

	"github.com/urfave/cli"
)

var defaultRequest = httpclient.Request{"GET", "www.baidu.com", nil}

func Client(c *cli.Context) error {
	var (
		keyProtocol = "protocol"
		keyUrl      = "url"
		keyCaseName = "casename"
	)

	protocol := c.String(keyProtocol)
	url := c.String(keyUrl)
	caseName := c.String(keyCaseName)
	log.Printf("protocol:%s, url:%s, case_name:%s", protocol, url, caseName)

	client := GetClientByProtocal(protocol)
	request, err := GetRequest(url, caseName)
	if err != nil {
		log.Printf("get request fail, error_msg:%s", err.Error())
		return err
	}

	body, err := client.Do(request)
	if err != nil {
		log.Printf("send request fail, error_msg:%s", err.Error())
		return err
	}
	log.Printf("body:%s", string(body))
	return nil
}

func GetClientByProtocal(p string) *httpclient.HttpClient {
	return httpclient.NewHttpClient(false)
}

func GetRequest(url, caseName string) (*httpclient.Request, error) {
	if url != "" {
		return MakeCaseWithUrl(url)
	}

	if caseName != "" {
		return MakeCaseWithCaseName(caseName)
	}

	return &defaultRequest, nil
}

func MakeCaseWithCaseName(caseName string) (*httpclient.Request, error) {
	return nil, nil
}

func MakeCaseWithUrl(url string) (*httpclient.Request, error) {
	return &httpclient.Request{
		Method: "GET",
		Url:    url,
	}, nil
}
