package global

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"

	"github.com/glstr/futty_golang/service/data/sqldb"
)

type HttpServerConfig struct {
	ServerAddr string `json:"server_addr"`
	DebugAddr  string `json:"debug_addr"`
}

type LogConfig struct {
	LogPath string `json:"log_path"`
}

type SqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"db_name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	LogConf        LogConfig        `json:"log_conf"`
	HttpServerConf HttpServerConfig `json:"http_server_conf"`
	SqlConf        SqlConfig        `json:"sql_conf"`

	// global client
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

var GCliResource ClientResource

type ClientResource struct {
	SqlDB *sql.DB
}

func (r *ClientResource) Init(conf *Config) error {
	sqlConf := conf.SqlConf
	option := sqldb.DBOption{
		Type:     sqldb.MysqlDBType,
		Host:     sqlConf.Host,
		Port:     sqlConf.Port,
		DBName:   sqlConf.DBName,
		User:     sqlConf.User,
		Password: sqlConf.Password,
	}
	db, err := sqldb.GetDB(&option)
	if err != nil {
		return err
	}

	GCliResource.SqlDB = db
	return nil
}
