//   A[0] = 10    A[1] = 2    A[2] = 5
//   A[3] = 1     A[4] = 8    A[5] = 20

// P < Q < R
// [10, 2, 5]
// [2, 5, 1]
// [5, 1, 8]
// [1, 8 ,20]
// [10, 5, 8]

// A[P] + A[Q] > A[R]
// A[Q] + A[R] > A[P]
// A[R] + A[P] > A[Q]

// [10, 2, 5, 1, 8, 20]

// [1, 2, 5, 8, 10, 20]

package solution

import "sort"

func Solution(A []int) int {
	if len(A) < 3 {
		return 0
	}

	sort.IntSlice(A).Sort()

	for i := 0; i < len(A)-2; i++ {
		if A[i]+A[i+1] > A[i+2] {
			return 1
		}
	}

	return 0
}
