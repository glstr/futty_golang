package utils

import "testing"

func TestGetTargetFilesFromDir(t *testing.T) {
	result := GetTargetFilesFromDir("/media/pi/glstr/download/download", "mp4")
	t.Logf("result:%v", result)
}
