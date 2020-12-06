package stack

// 普通的栈
// Abstract Data Type
type ADT interface {
	Push(data interface{})
	Top() interface{}
	Pop() interface{}
	IsEmpty() bool
}

type Stack struct {
	Data []interface{}
	Top_ int
	Base int
}

func Init() *Stack {
	S := &Stack{
		Data: make([]interface{}, 0, 100),
		Top_: 0,
		Base: 0,
	}

	return S
}

func (s *Stack) Push(data interface{}) {
	s.Data = append(s.Data, data)
	s.Top_++
}

func (s *Stack) Top() interface{} {
	return s.Data[s.Top_-1]
}

func (s *Stack) Pop() interface{} {
	data := s.Data[s.Top_-1]

	s.Data = s.Data[0 : len(s.Data)-1]
	s.Top_--

	return data
}

func (s *Stack) IsEmpty() bool {
	return s.Top_ == s.Base
}


