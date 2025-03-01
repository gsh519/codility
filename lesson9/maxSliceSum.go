package solution

func Solution(A []int) int {
	max := A[0]
	sum := 0

	for i := 0; i < len(A); i++ {
		sum += A[i]

		if max < sum {
			max = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return max
}
