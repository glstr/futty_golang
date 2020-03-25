package model

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	ItemType = "href"
)

var (
	ErrRequestFail = errors.New("request fail")
)

type HTMLParser struct {
}

type Targets struct {
	Element string
	Attr    string
}

func (p *HTMLParser) GetTargets(rawurl string, target string) ([]string, error) {
	var result []string
	u, err := url.Parse(rawurl)
	if err != nil {
		return result, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		return result, err
	}
	//res, err := http.Get(url)
	res, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return result, ErrRequestFail
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return result, err
	}

	doc.Find(target).Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		src, ok := s.Attr("src")
		if ok {
			src = u.Scheme + "://" + u.Host + src
			result = append(result, src)
		}
	})
	return result, nil
}
