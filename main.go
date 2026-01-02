package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/glstr/futty_golang/global"
	"github.com/glstr/futty_golang/httpserver"
	"github.com/glstr/futty_golang/logger"
)

func initConf() error {
	confPath := "./conf/snow.conf"
	return global.GConfig.Load(confPath)
}

func initLog() error {
	logPath := global.GConfig.LogConf.LogPath
	option := &logger.LogOption{
		LogPath: logPath,
	}
	return logger.InitLogger(option)
}

func initClientResource() error {
	return global.GCliResource.Init(&global.GConfig)
}

func initPprof() error {
	go func() {
		err := http.ListenAndServe(":8764", nil)
		if err != nil {
			panic(err)
		}
	}()
	return nil
}

func InitModule() error {
	err := initConf()
	if err != nil {
		return err
	}

	err = initLog()
	if err != nil {
		return err
	}

	err = initClientResource()
	if err != nil {
		return err
	}

	initPprof()
	return nil
}

func main() {
	//init module,including conf log pprof and so on.
	err := InitModule()
	if err != nil {
		panic(err)
	}

	logger.Notice("start http service")
	err = httpserver.StartHttpServer()
	if err != nil {
		panic(err)
	}
}
