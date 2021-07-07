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
