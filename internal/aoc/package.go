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

func Peek[T any](stack Stack[T]) T {
	n := len(stack)
	return stack[n-1]
}

func Pop[T any](stack Stack[T]) (Stack[T], T) {
	n := len(stack)
	return stack[:n-1], stack[n-1]
}

func Push[T any](stack Stack[T], item T) Stack[T] {
	return append(stack, item)
}
