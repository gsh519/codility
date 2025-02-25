package solution

func Solution(S string) int {
	if len(S) == 0 {
		return 1
	}

	count := 0

	for _, c := range S {
		if c == '(' {
			count++
			continue
		}

		if count == 0 {
			return 0
		}

		count--
	}

	if count == 0 {
		return 1
	}

	return 0
}
