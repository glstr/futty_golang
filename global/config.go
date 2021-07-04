package global

import (
	"encoding/json"
	"io/ioutil"
)

type HttpServerConfig struct {
	ServerAddr string `json:"server_addr"`
	DebugAddr  string `json:"debug_addr"`
}

type LogConfig struct {
	LogPath string `json:"log_path"`
}

type Config struct {
	LogConf        LogConfig `json:"log_conf"`
	HttpServerConf string    `json:"http_server_conf"`
}

var GConfig Config

func (c *Config) Load(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(content), c)
	if err != nil {
		return err
	}
	return nil
}
