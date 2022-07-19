package service

type CmdService interface {
	Exec(method string, args ...string) (string, error)
}

var cmdSers = map[string]CmdService{
	"test": NewPyCmdService(),
}

func GetCmdService() CmdService {
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

	return "", ErrNotFoundCmd
}
