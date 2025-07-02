package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got, want := Sum(numbers), 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got, want := SumAll([]int{1, 2}, []int{0, 9}), []int{3, 9}
	assertSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sus of some slices", func(t * testing.T) {
		got, want := SumAllTails([]int{1, 2}, []int{0, 9}), []int{2, 9}
		assertSums(t, got, want)
	})
	
	t.Run("safely sum empty slices", func(t *testing.T) { 
		got, want := SumAllTails([]int{}, []int{3, 4, 5}), []int{0, 9}
		assertSums(t, got, want)
	})
}

func assertSums(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want % v", got, want)
	}
}
