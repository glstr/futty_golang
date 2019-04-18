package howuse

import (
	"bytes"
	"log"
	"os/exec"
)

//ExecCmd provides method to use shell in go
func ExecCmd() string {
	cmd := exec.Command("ls", "-al")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	log.Printf("res:%s", out.String())
	return out.String()
}

func ExecPythonCmd() string {
	filePath := "/Users/pengbaojiang/pengbaojiang/code/gosrc/futty_golang/bin/replace_replay.py"
	cmd := exec.Command("python", filePath, "edit", "{\"file_key\":\"rtmp.liveshow.lss-user.baidubce.com/live/stream_bduid_809480822_1265438581/recording_20171229104952.m3u8\", \"begin\": 30, \"end\": 80}")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	log.Printf("res:%s", out.String())
	return out.String()
}
