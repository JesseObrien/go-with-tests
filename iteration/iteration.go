package iteration

import "strings"

// Repeat will repeat characters
func Repeat(characters string, times int) string {
	repeatSlice := make([]string, times)
	for i := 0; i < times; i++ {
		repeatSlice = append(repeatSlice, characters)
	}
	return strings.Join(repeatSlice, "")
}

// Compare two strings lexicographically
func Compare(a string, b string) int {
	if a == b {
		return 0
	}

	if a < b {
		return -1
	}

	return 1
}
