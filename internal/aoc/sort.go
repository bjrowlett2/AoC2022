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
