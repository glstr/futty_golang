package tool

import (
	"net/http"

	"github.com/glstr/futty_golang/global"
	"github.com/glstr/futty_golang/logger"
)

var (
	map2InitFunc = map[string]func() error{
		"conf":   initConf,
		"log":    initLog,
		"client": initClientResource,
		"pprof":  initPprof,
	}
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
	for _, initFunc := range map2InitFunc {
		err := initFunc()
		if err != nil {
			return err
		}
	}

	return nil
}

func StartHttpServer() error {
	//init module,including conf log pprof and so on.
	err := InitModule()
	if err != nil {
		panic(err)
	}

	logger.Notice("start http service")
	err = StartHttpServer()
	if err != nil {
		panic(err)
	}
	return nil
}
