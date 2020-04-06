package model

import (
	"net/url"
	"strings"
)

type UrlHelper struct {
	Url         string
	PathSection []string
	PurePath    string
}

func ParseUrl(rawurl string) (*UrlHelper, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	uh := &UrlHelper{
		Url: rawurl,
	}

	uh.PurePath = u.Path
	uh.PathSection = strings.Split(uh.PurePath, "/")
	return uh, nil
}

func (h *UrlHelper) GetPathSection() []string {
	return h.PathSection
}

func (h *UrlHelper) GetPurePath() string {
	return h.PurePath
}

func (h *UrlHelper) GetHost() string {
	return ""
}
