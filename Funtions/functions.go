package main

func addAndMultiply(a, b int) (int, int) {
	sum := a + b
	multiply := a * b
	return sum, multiply
}

func square(a int) (squa int) {
	squa = a * a
	return
}

// fucntions with variadic parameters (Unlimited args)

func sumAll(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}
