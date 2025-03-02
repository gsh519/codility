package solution

import "math"

func Solution(A []int) int {
	distincts := make(map[int]bool)
	count := 0

	for i := 0; i < len(A); i++ {
		if _, ok := distincts[int(math.Abs(float64(A[i])))]; !ok {
			distincts[int(math.Abs(float64(A[i])))] = true
			count++
		}
	}

	return count
}
