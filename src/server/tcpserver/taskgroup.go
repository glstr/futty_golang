package tcpserver

type Task struct {
	req *MessageRequest
	s   Session
}

type TaskGroup struct {
	workerNum uint32
}

func NewTaskGroup(workerNum uint32) *TaskGroup {
	return &TaskGroup{
		workerNum: workerNum,
	}
}

func (t *TaskGroup) Init() error {
	return nil
}

func (t *TaskGroup) Submit(task Task) error {
	return nil
}
