package calculate

func GetAverage(numbers []int, ch chan float64) {
	var sum = 0

	for _, e := range numbers {
		sum += e
	}

	ch <- float64(sum) / float64(len(numbers))
}

func GetMax(numbers []int, ch chan int) {
	var max = numbers[0]

	for _, e := range numbers {
		if e > max {
			max = e
		}
	}

	ch <- max
}
