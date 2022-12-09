package aoc

type Set[T comparable] map[T]bool

func (set *Set[T]) Add(item T) {
	if set == nil {
		return
	}

	s := *set
	s[item] = true
}
