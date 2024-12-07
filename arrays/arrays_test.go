package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		want := 15
		got := Sum(numbers)

		assertEquals(t, want, got, numbers)
	})

	t.Run("6 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		want := 21
		got := Sum(numbers)

		assertEquals(t, want, got, numbers)
	})
}

func assertEquals(t *testing.T, want int, got int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("want %d, got %d, given %v", want, got, numbers)
	}
}
