package main

import (
	"fmt"
)

func largestNumber(nums []int) string {
	n := len(nums)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return fmt.Sprintf("%d", nums[0])
	}
	temp := make([]string, n)
	r := 0
	for i := 0; i < n; i++ {
		temp[i] = fmt.Sprintf("%d", nums[i])
		r += nums[i]
	}
	if r == 0 {
		return "0"
	}
	quickSort(0, n-1, temp)
	fmt.Println(temp)
	str := ""
	for i := 0; i < n; i++ {
		str += temp[i]
	}
	return str
}

func quickSort(s, e int, nums []string) {
	if s >= e {
		return
	}
	s1 := s
	e1 := e - 1
	for s1 <= e1 {
		if comp(nums[e], nums[s1]) {
			s1++
			continue
		}
		if !comp(nums[e], nums[e1]) {
			e1--
			continue
		}
		nums[e1], nums[s1] = nums[s1], nums[e1]
	}
	nums[s1], nums[e] = nums[e], nums[s1]
	quickSort(s, s1-1, nums)
	quickSort(s1+1, e, nums)
}

func comp(a, b string) bool {
	if a == b {
		return false
	}
	n1 := len(a)
	n2 := len(b)
	p := 0
	for p < n1+n2 {
		if a[p%n1] < b[p%n2] {
			return true
		} else if a[p%n1] > b[p%n2] {
			return false
		}
		p++
	}
	return false
}

//  第二种方法
//type stringSlice []string
//
//func (ss stringSlice) Len() int {
//	return len(ss)
//}
//
//func (ss stringSlice) Swap(i, j int) {
//	ss[i], ss[j] = ss[j], ss[i]
//}
//
//func (ss stringSlice) Less(i, j int) bool {
//	// e.g., "210" should be greater than "102"
//	return ss[i]+ss[j] > ss[j]+ss[i]
//}
//
//func largestNumber(nums []int) string {
//	if len(nums) == 0 {
//		return ""
//	}
//	strs := make([]string, 0, len(nums))
//	for _, value := range nums {
//		strs = append(strs, strconv.Itoa(value))
//	}
//	sort.Sort(stringSlice(strs))
//	// Corner case
//	if strs[0][0] == '0' {
//		return "0"
//	}
//	return strings.Join(strs, "")
//}

func main() {
	a := []int{12, 121}
	fmt.Println(largestNumber(a))

}
