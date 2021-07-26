package service

import "github.com/glstr/futty_golang/model"

type CmdService interface {
	Exec(method string, args ...string) (string, error)
}

var cmdSers = map[int]CmdService{
	1: NewPyCmdService(),
}

func GetCmdService(serviceType int) CmdService {
	if ser, ok := cmdSers[serviceType]; ok {
		return ser
	}

	return NewPyCmdService()
}

type PyCmdService struct {
	cmd   string
	paths map[string]string
}

func NewPyCmdService() *PyCmdService {
	return &PyCmdService{
		cmd: "/usr/bin/python",
		paths: map[string]string{
			"test": "/home/pi/pengbaojiang/code/pythonsrc/main.py",
		},
	}
}

func (s *PyCmdService) Exec(method string, args ...string) (string, error) {
	if path, ok := s.paths[method]; ok {
		cmd := Cmd{
			Name: s.cmd,
			Path: path,
		}
		return cmd.Exec(args...)
	}

	return "", model.ErrNotFoundCmd
}
