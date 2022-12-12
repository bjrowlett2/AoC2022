package aoc

type Stack[T comparable] []T

func (stack *Stack[T]) Peek(item *T) bool {
	if stack == nil {
		return false
	}

	slice := *stack
	if len(slice) <= 0 {
		return false
	}

	n := len(slice)
	*item = slice[n-1]
	return true
}

func (stack *Stack[T]) Pop(item *T) bool {
	if stack == nil {
		return false
	}

	slice := *stack
	if len(slice) <= 0 {
		return false
	}

	n := len(slice)
	*item = slice[n-1]
	*stack = slice[:n-1]
	return true
}

func (stack *Stack[T]) Push(item T) {
	*stack = append(*stack, item)
}
