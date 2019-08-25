package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	things := Repeat("things", 2)
	fmt.Println(things)
	// Output: thingsthings
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestCompare(t *testing.T) {
	t.Run("test a is equal to a produces 0", func(t *testing.T) {
		comparison := Compare("a", "a")
		expected := 0

		if comparison != expected {
			t.Errorf("expected '%d', got '%d'", expected, comparison)
		}
	})

	t.Run("test a larger than b produces -1", func(t *testing.T) {
		comparison := Compare("a", "b")
		expected := -1

		if comparison != expected {
			t.Errorf("expected '%d', got '%d'", expected, comparison)
		}
	})

	t.Run("test a lesser than b produces 1", func(t *testing.T) {
		comparison := Compare("b", "a")
		expected := 1

		if comparison != expected {
			t.Errorf("expected '%d', got '%d'", expected, comparison)
		}
	})
}

func ExampleCompare() {
	comparison := Compare("a", "a")
	fmt.Println(comparison)
	// Output: 0
}
