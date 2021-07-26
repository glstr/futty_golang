package service

import "testing"

func TestCmdExec(t *testing.T) {
	output, err := DefaultCmd.Exec()
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
	}
	t.Logf("output:%s", string(output))
}
