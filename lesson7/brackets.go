package solution

func Solution(S string) int {
	mapped := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	var expected []rune

	for _, v := range S {
		if v == '[' || v == '(' || v == '{' {
			expected = append(expected, mapped[v])
			continue
		}

		if len(expected) == 0 || v != expected[len(expected)-1] {
			return 0
		}

		expected = expected[:len(expected)-1]
	}

	if len(expected) == 0 {
		return 1
	}
	return 0
}
