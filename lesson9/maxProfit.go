package solution

func Solution(A []int) int {
	p := 0

	if len(A) == 0 {
		return p
	}

	min := A[0]

	for i := 1; i < len(A); i++ {
		if A[i] < min {
			min = A[i]
			continue
		}

		if min < A[i] {
			currentProfit := A[i] - min
			if p < currentProfit {
				p = currentProfit
			}
		}
	}

	return p
}
