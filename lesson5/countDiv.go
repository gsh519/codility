package solution

func Solution(A int, B int, K int) int {
	count := int(B/K) - int(A/K)
	if A%K == 0 {
		count++
	}

	return count
}
