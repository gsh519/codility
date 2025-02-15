package solution

import (
	"sort"
)

func Solution(A []int) int {
	sort.IntSlice(A).Sort()

	for i := 0; i < len(A); i++ {
		if i+1 == A[i] {
			continue
		}

		return 0
	}

	return 1
}
