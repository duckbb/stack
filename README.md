# stack
standard library container/list to build stack

##install
go get github.com/duckbb/stack
##quick start
```api
s := NewStack()

//push
s.Push(1)
s.Push(2)
s.Push(3)
s.Push(4)
s.Push(5)

//pop
v, ok := s.Pop()

//remove 
s.Remove(1)

//exitst
find := s.Exist(2)

//length
length := s.Len()

```

