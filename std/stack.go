package std

// Stack is simple array based implementation of stack which accepts items of type interface{}
// which will require appropriate casting while using.
type Stack struct {
	ar []interface{}
}

func NewStack() *Stack {
	return &Stack{
		ar: make([]interface{}, 0),
	}
}

func (s *Stack) Push(item interface{}) {
	s.ar = append(s.ar, item)
}

func (s *Stack) Pop() interface{} {
	if s.Size() == 0 {
		panic("pop on empty stack")
	}
	n := len(s.ar)
	last := s.ar[n-1]
	s.ar = s.ar[:n-1]
	return last
}

func (s *Stack) Peek() interface{} {
	if s.Size() == 0 {
		panic("peek on empty stack")
	}
	return s.ar[len(s.ar)-1]
}

func (s *Stack) Size() int {
	return len(s.ar)
}
