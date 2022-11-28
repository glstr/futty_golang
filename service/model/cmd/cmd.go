package cmd

import (
	"os/exec"
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

func ListFilesInDir(dir string, fileSuffix string) ([]string, error) {
	return nil, nil
}
