package stack

type Interface interface {
	Len() int
	Top() interface{}
	Push(interface{})
	Pop()
}

func Empty(s Interface) bool {
	return 0 == Len(s)
}

func Len(s Interface) int {
	return s.Len()
}

func Top(s Interface) interface{} {
	return s.Top()
}
func Push(s Interface, x interface{}) {
	s.Push(x)
}

func Pop(s Interface) {
	s.Pop()
}
