package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("give me 5", func(t *testing.T) {
		got := Add(2, 3)
		want := 5

		if got != want {
			t.Errorf("expected sum %d, got %d", want, got)
		}

	})
	t.Run("give me 6", func(t *testing.T) {

		got := Add(1, 5)
		want := 6

		if got != want {
			t.Errorf("expected sum %d, got %d", want, got)
		}

	})

}

func ExampleAdd() {
	sum := Add(3, 4)
	fmt.Println(sum)
	// Output: 5
}
