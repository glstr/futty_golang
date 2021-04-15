package model

import "container/list"

type Queue struct {
	l *list.List
}

func NewQueue() *Queue {
	return &Queue{
		l: list.New(),
	}
}

func (q *Queue) Push(v interface{}) error {
	q.l.PushBack(v)
	return nil
}

func (q *Queue) Pop() interface{} {
	ele := q.l.Front()
	if ele == nil {
		return nil
	}
	return q.l.Remove(ele)
}

func (q *Queue) Len() int {
	return q.l.Len()
}

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{
		l: list.New(),
	}
}

func (s *Stack) Push(v interface{}) error {
	s.l.PushBack(v)
	return nil
}

func (s *Stack) Pop() interface{} {
	ele := s.l.Back()
	if ele == nil {
		return nil
	}
	return s.l.Remove(ele)
}

func (s *Stack) Len() int {
	return s.l.Len()
}
