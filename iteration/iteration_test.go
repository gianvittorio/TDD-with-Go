package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated, expected := Repeat("a", 5), "aaaaa"
	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
