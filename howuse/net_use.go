package howuse

import (
	"fmt"
	"net"
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

//IPUse show usage of IP
func IPUse() {
	ipStr := "127.0.0.1"
	ipStr = "10.10.10.10"
	ip := net.ParseIP(ipStr)
	ip = ip.To4()
	fmt.Printf("%v", ip)

	for _, b := range ip {
		fmt.Printf("%v\n", b)
	}

	if ip.IsLoopback() {
		fmt.Printf("it is a loopback address")
	}
}
