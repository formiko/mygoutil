package mygoutil

type StackInt []int

func (stack *StackInt) empty() bool {
	return 0 == len(*stack)
}

func (stack *StackInt) top() int {
	return (*stack)[len(*stack)-1]
}

func (stack *StackInt) push(v int) {
	*stack = append(*stack, v)
}

func (stack *StackInt) pop() {
	*stack = (*stack)[:len(*stack)-1]
}