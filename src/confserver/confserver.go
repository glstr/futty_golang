package confserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"utils"

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
	g.POST("/get_conf", s.getConf)
	g.POST("/update_conf", s.updateConf)
}

func (s *ConfServer) ServiceInit(confFile string) error {
	s.ConfManager = newConfManager(confFile)
	return nil
}

func (s *ConfServer) getConf(c *gin.Context) {
	ct := utils.NewContext()
	logBuffer := ct.LogBuffer
	logger := ct.Logger
	logBuffer.WriteLog("[body:%v]", c.Request.Body)

	defer func() {
		if err, ok := recover().(error); ok {
			logBuffer.WriteLog("[error_msg:%s]", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error_code": 2,
				"error_msg":  err.Error(),
			})
		}
		logger.Info(logBuffer.String())
	}()

	var getConfReq struct {
		ConfId string `json:"conf_id"`
	}

	if err := c.ShouldBindJSON(&getConfReq); err != nil {
		logBuffer.WriteLog("[err_msg:%s]", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error_msg": err.Error()})
		return
	}

	logBuffer.WriteLog("[confid:%s]", getConfReq.ConfId)
	conf := s.ConfManager.getConf(getConfReq.ConfId)
	if conf == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error_msg": "no conf"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"conf": conf})
	return
}

func (s *ConfServer) updateConf(c *gin.Context) {
	ct := utils.NewContext()
	logBuffer := ct.LogBuffer
	logger := ct.Logger

	defer func() {
		if err, ok := recover().(error); ok {
			logBuffer.WriteLog("[error_msg:%s]", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error_code": 2,
				"error_msg":  err.Error(),
			})
		}
		logger.Info(logBuffer.String())
	}()

	var updateConfReq struct {
		ConfId string `json:"conf_id"`
		Conf   string `json:"conf"`
	}

	if err := c.ShouldBindJSON(&updateConfReq); err != nil {
		logBuffer.WriteLog("[err_msg:%s]", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error_msg": err.Error()})
		return
	}

	ok := s.ConfManager.updateConf(updateConfReq.ConfId, updateConfReq.Conf)
	logBuffer.WriteLog("[update conf]")
	if ok {
		//persistent
		logBuffer.WriteLog("[persist conf]")
		err := s.ConfManager.persistConf()
		if err != nil {
			logBuffer.WriteLog("[err_msg:%s]", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error_msg": err.Error()})
			return
		}
	} else {
		errMsg := "update conf fail"
		logBuffer.WriteLog("[err_msg:%s]", errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error_msg": errMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error_code": 0})
	return
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
	Confs    map[string]interface{}
	mutex    sync.Mutex
	mutexPes sync.Mutex
}

type ConfFormat struct {
	version    string                 `json:"version"`
	Confs      map[string]interface{} `json:"confs"`
	UpdateTime string                 `json:"update_time"`
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
		confStr, ok := conf.(string)
		if ok {
			return confStr
		}
	}
	return InvalidConf
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
	confFile := "./conf/confserver.conf"
	f, err := os.Open(confFile)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(string(json))
	if err != nil {
		return err
	}
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
	err = json.Unmarshal(confContent, &confFormat)
	if err != nil {
		return err
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Confs = confFormat.Confs
	return nil
}

type confSubscribersManager struct {
	ConfSubsList map[string][]string
	Mutex        sync.Mutex
}

func (cfm *confSubscribersManager) subscribe(confId string,
	callbackAddr string) (string, error) {
	cfm.Mutex.Lock()
	defer cfm.Mutex.Unlock()
	cbAddrList, ok := cfm.ConfSubsList[confId]
	if ok {
		cbAddrList = append(cbAddrList, confId)
	} else {
		var addLists []string
		addLists = append(addLists, confId)
		cfm.ConfSubsList[confId] = addLists
	}
	return "", nil
}

func (cfm *confSubscribersManager) unsubscribe(subscribeId int64, confId string) error {
	cfm.Mutex.Lock()
	defer cfm.Mutex.Unlock()
	cbAddrList, ok := cfm.ConfSubsList[confId]
	if ok {
	}
	return nil
}
