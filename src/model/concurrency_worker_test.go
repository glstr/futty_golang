package model

import (
	"math/rand"
	"testing"
)

func MakeTask() []TaskHandler {
	var tasks []TaskHandler
	for i := 0; i < 10; i++ {
		task := func() *Result {
			return &Result{
				Num: rand.Int63n(100),
			}
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func TestDoTaskInGroup(t *testing.T) {
	var cw ConcurrencyWoker
	tasks := MakeTask()
	r, err := cw.DoTaskInGroup(tasks)
	if err != nil {
		t.Errorf("error_msg:%s", err.Error())
		return
	}
	t.Logf("result:%v", r)
}
