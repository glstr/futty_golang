package cmdhandler

import (
	"encoding/json"
	"errors"
	"os/exec"
)

var (
	ErrNotFoundCmd = errors.New("not found cmd")
)

type Cmd struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (c *Cmd) Exec(params ...string) (string, error) {
	var args []string
	args = append(args, c.Path)
	args = append(args, params...)
	cmd := exec.Command(c.Name, args...)
	res, err := cmd.Output()
	if err != nil {
		return string(res), err
	}
	return string(res), nil
}

var DefaultCmd Cmd = Cmd{"ls", "ls"}
var pythonTest Cmd = Cmd{"/usr/bin/python", "/home/pi/pengbaojiang/code/pythonsrc/main.py"}
var DefaultHandler *CmdHandler = NewCmdHander(map[string]Cmd{
	"ls":         DefaultCmd,
	"pythontest": pythonTest,
})

type CmdHandler struct {
	CmdGroup map[string]Cmd
}

func NewCmdHander(cmds map[string]Cmd) *CmdHandler {
	return &CmdHandler{
		CmdGroup: cmds,
	}
}

func (h *CmdHandler) Execute(cmdName string, params ...string) (string, error) {
	if cmd, ok := h.CmdGroup[cmdName]; ok {
		return cmd.Exec(params...)
	}
	return "", ErrNotFoundCmd
}

func (h *CmdHandler) String() string {
	output, err := json.Marshal(h.CmdGroup)
	if err != nil {
		return ""
	}
	return string(output)
}
