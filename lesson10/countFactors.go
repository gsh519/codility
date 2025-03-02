package solution

import "math"

func Solution(N int) int {
	count := 0
	sqrtN := int(math.Sqrt(float64(N)))

	for i := 1; i <= sqrtN; i++ {
		if N%i == 0 {
			count += 2
			if i*i == N {
				count--
			}
		}
	}

	return count
}
