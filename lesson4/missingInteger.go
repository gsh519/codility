package solution

import "sort"

func Solution(A []int) int {
	smallest := 1

	sort.IntSlice(A).Sort()

	for _, n := range A {
		if smallest < n {
			return smallest
		}
		if smallest == n {
			smallest++
		}
	}

	return smallest
}
