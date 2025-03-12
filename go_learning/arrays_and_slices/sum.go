package main

func Sum(numbers []int) (sum int) {
	for _, i := range numbers {
		sum += i
	}
	return
}

func SumAll(numbersToSum ...[]int) (sumall []int) {
	for _, numbers := range numbersToSum {
		sumall = append(sumall, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sumalltails []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sumalltails = append(sumalltails, 0)
			continue
		}
		tail := numbers[1:]
		sumalltails = append(sumalltails, Sum(tail))
	}
	return
}
