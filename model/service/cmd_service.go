package service

type CmdService interface{}

type PyCmdService struct {
	cmd     string
	methods map[string]string
}

func NewPyCmdService() *PyCmdService {
	return &PyCmdService{
		cmd: "/usr/bin/python",
		methods: map[string]string{
			"test": "/home/pi/pengbaojiang/code/pythonsrc/main.py",
		},
	}
}

func (s *PyCmdService) Exec(method string, args ...string) (string, error) {
	return "", nil
}
