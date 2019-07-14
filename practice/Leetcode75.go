package main

func sortColors(nums []int) {
	var p0, p1, p2 int
	for _, num := range nums {
		nums[p2] = 2
		p2++
		if num < 2 {
			nums[p1] = 1
			p1++
		}
		if num < 1 {
			nums[p0] = 0
			p0++
		}
	}
}

func main() {
	t := []int{2, 0, 2, 1, 1, 0}
	sortColors(t)
}
