package euler

func Problem1(start, finish int) int {
	var sum int
	for ; start < finish; start++ {
		if start%3 == 0 {
			sum = sum + start
		} else if start%5 == 0 {
			sum = sum + start
		}
	}
	return sum
}
