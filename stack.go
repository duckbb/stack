package stack

import (
	"container/list"
	"sync"
)

type Stack struct {
	lock sync.Mutex
	data *list.List
}

func NewStack() *Stack {
	return &Stack{data: list.New()}
}

//push
func (s *Stack) Push(v interface{}) interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.data.PushBack(v).Value
}

//pop
func (s *Stack) Pop() (interface{}, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	ele := s.data.Back()
	if ele == nil {
		return nil, false
	}
	return s.data.Remove(ele), true
}

//len
func (s *Stack) Len() int {
	return s.data.Len()
}

//get
func (s *Stack) Exist(value interface{}) bool {
	if s.Len() == 0 {
		return false
	}
	elem := s.data.Front()
	for elem != nil {
		if elem.Value == value {
			return true
		}
		elem = elem.Next()
	}
	return false
}

//rm one element
func (s *Stack) Remove(value interface{}) bool {
	if s.Len() == 0 {
		return false
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	var delElem *list.Element
	elem := s.data.Front()
	for elem != nil {
		if elem.Value == value {
			delElem = elem
			break
		}
		elem = elem.Next()
	}
	if delElem != nil {
		s.data.Remove(delElem)
		return true
	}
	return false
}
