package main

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, 0)
	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := make([]int, 0)

	for _, nums := range numbersToSum {
		sums = append(sums, nums[1:]...)
	}

	return sums
}
