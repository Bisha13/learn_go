package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	want := "aaaaa"
	got := Repeat("a")
	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
}
