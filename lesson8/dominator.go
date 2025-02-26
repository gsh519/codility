package solution

func Solution(A []int) int {
	half := len(A) / 2
	mapped := make(map[int]int)

	for i := 0; i < len(A); i++ {
		if _, ok := mapped[A[i]]; ok {
			mapped[A[i]]++
		} else {
			mapped[A[i]] = 1
		}

		if half < mapped[A[i]] {
			return i
		}
	}

	return -1
}
