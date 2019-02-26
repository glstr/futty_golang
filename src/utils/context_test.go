package utils

import "testing"

func init() {
	Load("conf/rdguard_go.conf")
	LogInit()
}

func TestWriteLog(t *testing.T) {
	c := NewContext()
	c.LogBuffer.WriteLog("[test:%s][logid:%d]", "hello_world", c.Logid)
	c.Logger.Info(c.LogBuffer.String())
}
