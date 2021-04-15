package model

import (
	"bytes"
	"client/httpclient"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	ItemType = "href"
)

type HTMLParser struct {
	client *httpclient.HttpClient
}

func NewHTMLParser() *HTMLParser {
	return &HTMLParser{
		client: httpclient.NewHttpClient(false),
	}
}

type Target struct {
	RawUrl  string
	Element string
	Attr    string
}

func (p *HTMLParser) GetTargetFromUrl(target *Target) ([]string, error) {
	var result []string
	baseUrl, err := url.Parse(target.RawUrl)
	if err != nil {
		return result, err
	}

	req := &httpclient.Request{
		Method: "GET",
		Url:    target.RawUrl,
	}

	body, err := p.client.Do(req)
	if err != nil {
		return result, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return result, err
	}

	doc.Find(target.Element).Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		srcUrl, ok := s.Attr(target.Attr)
		if ok {
			u, err := url.Parse(srcUrl)
			if err != nil {
				return
			}

			if u.IsAbs() {
				result = append(result, srcUrl)
			} else {
				abUrl := baseUrl.ResolveReference(u).String()
				result = append(result, abUrl)
			}
		}
	})

	return result, nil
}
