package solution

import "math"

func Solution(N int) int {
	min := 2 * (1 + N)
	sqrtN := int(math.Sqrt(float64(N)))

	for i := 2; i <= sqrtN; i++ {
		if N%i == 0 {
			a := i
			b := N / i

			r := 2 * (a + b)

			if r < min {
				min = r
			}
		}
	}

	return min
}
