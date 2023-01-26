package surface_test

import (
	"prac_go/surface"
	"testing"
)

func TestRect(t *testing.T) {
	got := surface.Rect(2, 4)
	want := 8

	if got != want {
		t.Errorf("expected surface.Rect() = %d, but got %d", want, got)
	}
}
