package model

import (
	"errors"
	"log"
	"message"
	"net"
	"runtime"
	"sync"
	"time"
)

var (
	statusOK     = 0
	statusClosed = 1
)

type Task interface {
	Work(c net.Conn) error
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
	status int
	cancel chan struct{}

	taskStream chan Task
	timeOut    time.Duration
	addr       string
	key        string
}

func (c *Conner) NewConner(addr string, taskStream chan Task,
	cancel chan struct{}) *Conner {
	return &Conner{
		cancel:     cancel,
		taskStream: taskStream,
		addr:       addr,
	}
}

func (c *Conner) ConnWork() {
	defer func() {
		if err, ok := recover().(error); ok {
			log.Printf("addr:%s, key:%s, error_msg:%s",
				c.addr, c.key, err.Error())
			var buf []byte
			n := runtime.Stack(buf, false)
			log.Printf("trace:%v", buf[:n])
		}

		if c.status != statusClosed {
			c.Close()
			return
		}
	}()

	for {
		select {
		case task := <-c.taskStream:
			if c.status != statusOK {
				err := c.makeConn()
				if err != nil {
					log.Printf("addr:%s, key:%s, error_msg:%s",
						c.addr, c.key, err.Error())
					return
				}
			}

			err := c.handleTask(task)
			if err != nil {
				c.Close()
				return
			}
		case <-c.cancel:
			c.Close()
			return
		case <-time.After(c.timeOut):
			c.Close()
			return
		}
	}
}

func (c *Conner) makeConn() error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	c.conn = conn
	c.status = statusOK
	return nil
}

func (c *Conner) handleTask(t Task) error {
	return t.Work(c.conn)
}

func (c *Conner) Close() error {
	c.conn.Close()
	c.status = statusClosed
	return nil
}

type Worker struct {
	key      int32
	connMu   sync.RWMutex
	connPool map[string]net.Conn
	connMax  int

	addr string

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
	WorkersNum int
	addrs      []string

	workersMu sync.RWMutex
	workers   map[string]*Worker
}

func NewWorkerGroup(addrs []string) *WorkerGroup {
	return &WorkerGroup{
		WorkersNum: len(addrs),
		addrs:      addrs,
	}
}

func (wg *WorkerGroup) Init() {
}

func (wg *WorkerGroup) Start() {
}

func (wg *WorkerGroup) Close() {
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
