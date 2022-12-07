package aoc

import "sort"

func SortInt64(slice []int64) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func SortInt64Desc(slice []int64) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[j] < slice[i]
	})
}

type Stack[T any] []T

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
