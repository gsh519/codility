package solution

import "sort"

func Solution(A []int) int {
	sort.IntSlice(A).Sort()

	x := A[len(A)-1] * A[len(A)-2] * A[len(A)-3]

	y := 0
	if A[0] < 0 && A[1] < 0 {
		y = A[0] * A[1] * A[len(A)-1]
	}

	if x < y {
		return y
	}

	return x
}
