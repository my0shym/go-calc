package calc

import "testing"

func TestMax(t *testing.T) {
	want := 2
	got := Max(1, 2)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
