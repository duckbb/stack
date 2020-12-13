package stack

import (
	"testing"
)

//test push
func TestAdd(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	expectNum := 5
	if s.Len() != expectNum {
		t.Error("add data, but data length is not equal!")
	}
}

// test pop
func TestPush(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	popList := []interface{}{5, 4, 3, 2, 1}
	for _, v := range popList {
		popdata, ok := s.Pop()
		if !ok || popdata != v {
			t.Errorf("pop data is:%v,but list data is:%v\n", popdata, v)
		}
	}
}

//test remove
func TestRemove(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Remove(4)
	expect := 4
	if s.Len() != expect {
		t.Errorf("remove one elme, but length is:%v,expect:%v", s.Len(), expect)
	}
}

//test Exist
func TestExist(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	cases := []struct {
		find   int
		expect bool
	}{
		{1, true},
		{2, true},
		{103, false},
		{4, true},
		{100, false},
	}

	for _, v := range cases {
		real := s.Exist(v.find)
		if real != v.expect {
			t.Errorf("find value:%v, expect:%v,actually is:%v\n", v.find, v.expect, real)
		}
	}
}
