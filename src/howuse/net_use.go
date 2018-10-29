package howuse

import (
	"fmt"
	"net/url"
)

func UserUrlParse() {
	eUrl := "https://golang.org/pkg/net/url/?hello=1&nihao=中国"
	res, err := url.Parse(eUrl)
	if !CheckError(err) {
		return
	}

	fmt.Printf("url:%v\n", res)
	fmt.Printf("host:%s\n", res.Host)
	fmt.Printf("path:%s\n", res.Path)
	fmt.Printf("query:%s\n", res.RawQuery)
}

func PostFile() {

}
