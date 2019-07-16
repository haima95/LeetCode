package main

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	i := 1
	j := 2
	for j < len(nums) {
		if nums[i] == nums[j] && nums[i-1] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}
	}
	return i + 1
}

func main() {
}
