package confserver

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

const InvalidConf = ""

type ConfServer struct {
	Router      *gin.Engine
	ConfManager *confManager
}

func NewConfServer(r *gin.Engine) *ConfServer {
	return &ConfServer{
		Router: r,
	}
}

func (s *ConfServer) LoadRouter() {
	g := s.Router.Group("confserver")
	g.GET("/confserver/get_conf", s.getConf)
}

func (s *ConfServer) ServiceInit(confFile string) {

}

func (s *ConfServer) getConf(c *gin.Context) {

}

func (s *ConfServer) updateConf(c *gin.Context) {

}

func (s *ConfServer) addConf(c *gin.Context) {

}

func (s *ConfServer) delConf(c *gin.Context) {

}

func (s *ConfServer) subscribeConf(c *gin.Context) {

}

func (s *ConfServer) unsubscribeConf(c *gin.Context) {

}

type confManager struct {
	Confs    map[string]string
	mutex    sync.Mutex
	mutexPes sync.Mutex
}

type ConfFormat struct {
	version    string
	Confs      map[string]string
	UpdateTime string
}

func newConfManager(confFile string) *confManager {
	ConfManager := &confManager{}
	err := ConfManager.loadConf(confFile)
	if err != nil {
		return nil
	} else {
		return ConfManager
	}
}

func (s *confManager) getConf(confId string) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	conf, ok := s.Confs[confId]
	if ok {
		return conf
	} else {
		return InvalidConf
	}
}

func (s *confManager) updateConf(confId string, newConf string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if newConf != InvalidConf {
		s.Confs[confId] = newConf
		return true
	} else {
		return false
	}
}

func (s *confManager) addConf(confId string, newConf string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.Confs[confId]
	if !ok {
		s.Confs[confId] = newConf
		return true
	}
	return false
}

func (s *confManager) delConf(confId string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.Confs, confId)
	return true
}

func (s *confManager) persistConf() error {
	var confFormat ConfFormat
	s.mutex.Lock()
	confFormat.Confs = s.Confs
	s.mutex.Unlock()

	confFormat.version = "test"
	confFormat.UpdateTime = "test time"
	json, err := json.Marshal(&confFormat)
	if err != nil {
		return err
	}

	s.mutexPes.Lock()
	defer s.mutexPes.Unlock()
	confFile := "conf/server.conf"
	f, err := os.Open(confFile)
	if err != nil {
		return err
	}
	f.WriteString(string(json))
	f.Close()
	return nil
}

func (s *confManager) loadConf(confFile string) error {
	f, err := os.Open(confFile)
	if err != nil {
		return err
	}
	confContent, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	var confFormat ConfFormat
	err = json.Unmarshal(confContent, confFormat)
	if err != nil {
		return err
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Confs = confFormat.Confs
	return nil
}

type confSubscribersManager struct {
}

func (cfm *confSubscribersManager) subscribe(confId string, callbackAddr string) error {
	return nil
}

func (cfm *confSubscribersManager) unsubscribe(subscribeId int64, confId string) error {
	return nil
}
