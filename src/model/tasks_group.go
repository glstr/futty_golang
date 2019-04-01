package model

import "fmt"

type TaskWorker struct {
	workList    chan work
	done        chan struct{}
	workDoneNum int64
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
			tw.workDoneNum++
		}
	}
}

type TasksGroup struct {
	workers        []*TaskWorker
	workList       chan work
	done           chan struct{}
	workListLength int64
	workersNum     int64
}

type work func()

func NewTasksGroup(wll, wn int64) *TasksGroup {
	return &TasksGroup{
		workListLength: wll,
		workersNum:     wn,
		done:           make(chan struct{}),
		workList:       make(chan work, wn),
	}
}

//Add provides operation to add work to worklist
func (t *TasksGroup) Add(w work) {
	t.workList <- w
}

//Cancel provides cancel operation for task group
func (t *TasksGroup) Cancel() {
	close(t.done)
}

//Stats provides stat info
func (t *TasksGroup) Stats() {
	for index, worker := range t.workers {
		fmt.Printf("index:%d, workDone:%d\n", index, worker.workDoneNum)
	}
}

//Start provides start task group loop
func (t *TasksGroup) Start() {
	//init workers
	for i := int64(0); i < t.workersNum; i++ {
		w := NewTaskWorker(t.workList, t.done)
		t.workers = append(t.workers, w)
	}

	//start loop
	for _, work := range t.workers {
		go work.Work()
	}
}
