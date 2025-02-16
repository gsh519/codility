package solution

func Solution(A []int) int {
	e := 0
	pair := 0

	for _, n := range A {
		if n == 0 {
			e++
		} else {
			pair += e
		}
	}

	if pair > 1000000000 {
		return -1
	}

	return pair
}
