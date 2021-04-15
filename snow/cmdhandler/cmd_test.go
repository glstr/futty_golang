package cmdhandler

import "testing"

func TestCmdExec(t *testing.T) {
	output, err := DefaultCmd.Exec()
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
	}
	t.Logf("output:%s", string(output))
}

func TestCmdHandler(t *testing.T) {
	output, err := DefaultHandler.Execute("ls")
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
	}
	t.Logf("output:%s", output)

	output, err = DefaultHandler.Execute("whoknowns")
	if err == nil {
		t.Errorf("expect err")
		return
	}
	t.Logf("output:%s", err.Error())
}
