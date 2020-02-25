package model

import (
	"log"
	"sync"
)

type Param struct{}
type Result struct {
	Num int64
}
type TaskHandler func() *Result
type ConcurrencyWoker struct {
}

func (w *ConcurrencyWoker) DoTaskInGroup(tasks []TaskHandler) (*Result, error) {
	var merger ResultMerger
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func() {
			result := task()
			err := merger.AddResult(result)
			if err != nil {
				log.Printf("error_mg:%s", err.Error())
			}
			log.Printf("result:%v", result)
			wg.Done()
		}()
	}
	wg.Wait()
	return merger.EndResult()
}

type ResultMerger struct {
	Results []*Result
	mutex   sync.Mutex
}

func (m *ResultMerger) AddResult(r *Result) error {
	m.mutex.Lock()
	m.Results = append(m.Results, r)
	m.mutex.Unlock()
	return nil
}

func (m *ResultMerger) EndResult() (*Result, error) {
	endResult := &Result{}
	m.mutex.Lock()
	for _, r := range m.Results {
		endResult.Num += r.Num
	}
	m.mutex.Unlock()
	return endResult, nil
}

func (w *ConcurrencyWoker) DoTaskInGroupWithChan(tasks []TaskHandler) (*Result, error) {
	return nil, nil
}
