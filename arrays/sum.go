package arrays

// Sum provides the sum total of an array of integers
func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum
}
