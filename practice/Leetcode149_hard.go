package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func slope(i, j int, points [][]int) string {
	p1 := points[i]
	p2 := points[j]
	if p1[0] == p2[0] {
		return "inf"
	}
	g := gcd((p1[1] - p2[1]), (p1[0] - p2[0]))
	num := fmt.Sprintf("%d", (p1[1]-p2[1])/g)
	den := fmt.Sprintf("%d", (p1[0]-p2[0])/g)
	return fmt.Sprintf("%d/%d", num, den)
}

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}
	ans := 0
	for i := 0; i < n; i++ {
		temp := make(map[string]int)
		count, same := 0, 0
		for j := i + 1; j < n; j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				same++
				continue
			}
			s := slope(i, j, points)
			if _, ok := temp[s]; ok {
				temp[s]++
			} else {
				temp[s] = 1
			}
			if count < temp[s] {
				count = temp[s]
			}
		}
		if ans < count+same {
			ans = count + same
		}
	}
	return ans + 1
}

func main() {

	temp := [][]int{{3, 1}, {12, 3}, {3, 1}, {-6, -1}}
	fmt.Println(maxPoints(temp))
}
