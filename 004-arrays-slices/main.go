package arraysslices

func Sum(nums []int) int {
	var sum int

	for _, n := range nums {
		sum += n
	}

	return sum
}

func SumAll(nums ...[]int) []int {
	var sums []int

	for _, n := range nums {
		sums = append(sums, Sum(n))
	}

	return sums
}
