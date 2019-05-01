package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		nums1,nums2 = nums2,nums1
		m,n = n,m
	}
	min := 0
	max := m
	halfLen := (m+n+1)/2
	for ;min <= max; {
		i := (min+max)/2
		j := halfLen - i
		if i < max && nums2[j-1] > nums1[i] {
			min = i+1
		}else if i > min && nums1[i-1] > nums2[j] {
			max = i-1
		}else {
			maxLeft := 0.0
			if i == 0 {
				maxLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = math.Max(float64(nums1[i-1]),float64(nums2[j-1]))
			}
			if (m+n)%2 == 1 {
				return maxLeft
			}

			minRight := 0.0
			if i == m {
				minRight = float64(nums2[j])
			}else if j == n {
				minRight = float64(nums1[i])
			}else {
				minRight = math.Min(float64(nums2[j]),float64(nums1[i]))
			}
			return (maxLeft+minRight)/2.0
		}
	}

	return 0.0
}

func main(){
	nums1 := []int{1,2}
	nums2 := []int{3,4,5}
	fmt.Println(findMedianSortedArrays(nums1,nums2))
}
