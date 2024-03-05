package main

func FindOddNumber(text []int) int {
	call := map[int]int{}

	for _, n := range text {
		call[n]++
	}

	for _, n := range text {
		if call[n]%2 != 0 {
			return n
		}
	}

	return -1
}
