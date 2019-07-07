package main

func plusOne(digits []int) []int {
	n := len(digits)
	temp := 1
	for i := n - 1; i > -1; i-- {
		temp = digits[i] + temp
		digits[i] = temp % 10
		temp = temp / 10
	}
	var result []int
	if temp > 0 {
		result = append(result, temp)
	}
	result = append(result, digits...)
	return result
}

func main() {
}
