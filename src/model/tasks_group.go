package model

type TaskWorker struct {
	workList chan work
	done     chan struct{}
}

func NewTaskWorker(wl chan work, d chan struct{}) *TaskWorker {
	return &TaskWorker{
		workList: wl,
		done:     d,
	}
}

func (tw *TaskWorker) Work() {
	for {
		select {
		case <-tw.done:
			return
		case work := <-tw.workList:
			work()
		}
	}
}

type TasksGroup struct {
	workers        []*TaskWorker
	workList       chan work
	workListLength int64
	workersNum     int64
}

type work func()

func NewTasksGroup(wll, wn int64) *TasksGroup {
	return &TasksGroup{}
}

func (t *TasksGroup) Add(w work) {
}

func (t *TasksGroup) Cancel() {
}

func (t *TasksGroup) Stats() {
}

func (t *TasksGroup) Start() {
}
