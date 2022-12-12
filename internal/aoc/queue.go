package aoc

type Queue[T comparable] []T

func (queue *Queue[T]) Peek(item *T) bool {
	if queue == nil {
		return false
	}

	slice := *queue
	if len(slice) <= 0 {
		return false
	}

	*item = slice[0]
	return true
}

func (queue *Queue[T]) Pop(item *T) bool {
	if queue == nil {
		return false
	}

	slice := *queue
	if len(slice) <= 0 {
		return false
	}

	*item = slice[0]
	*queue = slice[1:]
	return true
}

func (queue *Queue[T]) Push(item T) {
	*queue = append(*queue, item)
}

func (queue *Queue[T]) Contains(item T) bool {
	if queue == nil {
		return false
	}

	slice := *queue
	for _, value := range slice {
		if value == item {
			return true
		}
	}

	return false
}
