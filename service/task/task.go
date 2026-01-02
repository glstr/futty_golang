package task

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

type TaskState int

const (
	TaskStatInit   = 0
	TaskStatDone   = 1
	TaskStatFailed = 2
)

type TaskFunc func() error

type TaskResult struct {
	state     TaskState
	extraInfo []byte
	err       error
}

func (r *TaskResult) GetState() TaskState {
	return r.state
}

func (r *TaskResult) GetExtraInfo() []byte {
	return r.extraInfo
}

func (r *TaskResult) GetError() error {
	return r.err
}

type TaskService interface {
	SetTask(task TaskFunc, extraInfo []byte) (int32, error)
	GetTask(taskID int32) (*TaskResult, error)
}

type taskService struct {
	taskID      int32
	taskResults map[int32]*TaskResult
	sync.RWMutex
}

var (
	defaultTaskService = NewTaskService()
	ErrTaskNotFound    = errors.New("not found")
)

func GetTaskService() TaskService {
	return defaultTaskService
}

func NewTaskService() *taskService {
	return &taskService{
		taskID:      0,
		taskResults: make(map[int32]*TaskResult),
	}
}

func (s *taskService) SetTask(task TaskFunc, extraInfo []byte) (int32, error) {
	r := &TaskResult{
		state:     TaskStatInit,
		extraInfo: extraInfo,
	}
	taskID := s.getCurTaskID()
	s.setTask(taskID, r)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("程序捕获到异常: %v\n", r)
			}
		}()
		err := task()
		if err != nil {
			r.err = err
			r.state = TaskStatFailed
			return
		}
		r.state = TaskStatDone
	}()

	return taskID, nil
}

func (s *taskService) setTask(taskID int32, r *TaskResult) {
	defer s.Unlock()
	s.Lock()
	s.taskResults[taskID] = r
}

func (s *taskService) GetTask(taskID int32) (*TaskResult, error) {
	defer s.RUnlock()
	s.RLock()
	r, ok := s.taskResults[taskID]
	if ok {
		return r, nil
	}
	return nil, ErrTaskNotFound
}

func (s *taskService) getCurTaskID() int32 {
	return atomic.AddInt32(&s.taskID, 1)
}
