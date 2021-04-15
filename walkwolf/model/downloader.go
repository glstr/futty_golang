package model

import (
	"client/httpclient"
	"errors"
	"os"
	"strings"
)

var (
	ErrInvalidTask = errors.New("invalid task")
)

type Downloader struct {
	client *httpclient.HttpClient
}

func NewDownloader() *Downloader {
	return &Downloader{
		client: httpclient.NewHttpClient(false),
	}
}

type DownloadTask struct {
	URL  string
	Name string
	Dir  string
}

func (t *DownloadTask) path() string {
	filename := t.Name
	if filename == "" {
		filename = t.getFileNameFromUrl(t.URL)
	}

	if t.Dir != "" {
		return t.Dir + "/" + filename
	}

	return filename
}

func (t *DownloadTask) getFileNameFromUrl(Url string) string {
	strs := strings.Split(Url, "/")
	if len(strs) <= 0 {
		return "default"
	}
	return strs[len(strs)-1]
}

func (d *Downloader) RunTask(task *DownloadTask) error {
	if task == nil {
		return ErrInvalidTask
	}

	req := &httpclient.Request{
		Method: "GET",
		Url:    task.URL,
	}

	body, err := d.client.Do(req)
	if err != nil {
		return err
	}

	return d.OutputToFile(task.path(), body)
}

func (d *Downloader) OutputToFile(filePath string, body []byte) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	_, err = f.Write(body)
	return err
}
