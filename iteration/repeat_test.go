package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("give me aaaaa", func(t *testing.T) {
		want := "aaaaa"
		got := Repeat("a")
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}

	})
	t.Run("give me bbbbb", func(t *testing.T) {
		want := "bbbbb"
		got := Repeat("b")
		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}

	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
