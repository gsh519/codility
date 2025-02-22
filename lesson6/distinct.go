package solution

func Solution(A []int) int {
	d := make(map[int]bool)
	r := 0
	for _, a := range A {
		if _, v := d[a]; !v {
			d[a] = true
			r++
		}
	}

	return r
}
