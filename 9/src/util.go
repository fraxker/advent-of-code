package src

func PerdictNext(numbers []int) int {
	diff := make([]int, len(numbers)-1)
	var notZero bool
	for i := 0; i < len(numbers)-1; i++ {
		diff[i] = numbers[i+1] - numbers[i]
		if diff[i] != 0 {
			notZero = true
		}
	}
	if !notZero {
		return numbers[len(numbers)-1]
	}
	return PerdictNext(diff) + numbers[len(numbers)-1]
}

func PredictPrev(numbers []int) int {
	diff := make([]int, len(numbers)-1)
	var notZero bool
	for i := 0; i < len(numbers)-1; i++ {
		diff[i] = numbers[i+1] - numbers[i]
		if diff[i] != 0 {
			notZero = true
		}
	}
	if !notZero {
		return numbers[0]
	}
	return numbers[0] - PredictPrev(diff)
}