package model

import (
	"errors"
	"message"
	"net"
	"sync"
	"time"
)

type Task interface {
}

type MessageTask struct {
	Key     string
	Req     *message.Message
	Res     *message.Message
	Timeout time.Duration

	CallbackStream chan string
}

func (mt *MessageTask) Work(c net.Conn) error {
	return nil
}

type Conner struct {
	conn   net.Conn
	cancel chan struct{}
	done   chan string

	taskStream chan *Task
	timeOut    time.Duration
	addr       string
	key        string
}

func (c *Conner) NewConner(addr string, taskStream chan *Task,
	done chan string) *Conner {
	return &Conner{
		done:       done,
		taskStream: taskStream,
		addr:       addr,
	}
}

func (c *Conner) ConnWork() {
	for {
		select {
		case task := <-c.taskStream:
			err := c.handleTask(task)
			if err != nil {
				c.Close()
				return
			}
			return
		case <-c.cancel:
			c.Close()
			return
		case <-time.After(c.timeOut):
			c.Close()
			return
		}
	}
}

func (c *Conner) handleTask(t *Task) error {
	return nil
}

func (c *Conner) Close() error {
	c.conn.Close()
	return nil
}

type Worker struct {
	key      int32
	connMu   sync.RWMutex
	connPool map[string]net.Conn

	addr    string
	connMax int

	taskStream    chan *Task
	taskStreamMax int

	cancel chan struct{}
}

func NewWorker() *Worker {
	return &Worker{}
}

func (w *Worker) Init() error {
	return nil
}

func (w *Worker) Start() error {
	return nil
}

func (w *Worker) Update() error {
	for {
		select {
		case <-time.After(10 * time.Second):
			//delete invalid conn

			//if need, start new conn
			if len(w.connPool) < w.connMax && len(w.taskStream) > 0 {
			}
		}
	}
	return nil
}

func (w *Worker) Close() error {
	return nil
}

func (w *Worker) CommitTask(t *Task) error {
	select {
	case w.taskStream <- t:
		return nil
	default:
		return errors.New("task stream full")
	}
}

type WorkerGroup struct {
	WorkersNum int32

	workersMu sync.RWMutex
	workers   map[string]*Worker
	addrs     []string
}

func NewWorkerGroup() *WorkerGroup {

}

func (wg *WorkerGroup) CommitTask(t *Task) error {
	worker := wg.getWorker(t)
	if worker != nil {
		return errors.New("no invalid worker")
	}

	return worker.CommitTask(t)
}

func (wg *WorkerGroup) getWorker(t *Task) *Worker {
	return wg.workers[""]
}

func (wg *WorkerGroup) Cancel() error {
	return nil
}
