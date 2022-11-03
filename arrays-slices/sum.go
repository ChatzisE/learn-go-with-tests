package main

func Sum(numbers []int) int {
	sum := 0
	// _ (black identifier) if we don't to use the index i
	for _, number := range numbers {
		sum += number
	}
	return sum
}
