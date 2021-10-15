package arraysslices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}

	got := Sum(nums)
	want := 15

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleSum() {
	result := Sum([]int{1, 2, 3})
	fmt.Println(result)
	// Output: 6
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{1, 1, 1})
	want := []int{6, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSumAll() {
	result := SumAll([]int{1, 1, 1}, []int{2, 2, 2}, []int{3, 3})
	fmt.Println(result)
	// Output: [3 6 6]
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of filled slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 3, 1})
		want := []int{2, 4}

		checkSums(t, got, want)
	})

	t.Run("sately sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}

		checkSums(t, got, want)
	})
}

func ExampleSumAllTails() {
	result := SumAllTails([]int{1, 2}, []int{3, 3, 1})
	fmt.Println(result)
	// Output: [2 4]
}