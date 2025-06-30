package integers

import (
	"testing"
)

func TestAdder(t *testing.T) {
	sum, expected := Add(2, 2), 4
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}