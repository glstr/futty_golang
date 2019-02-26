package utils

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	LogPath string `json:"log_path"`
}

var Config config

func Load(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(content), &Config)
	if err != nil {
		panic(err)
	}
}
