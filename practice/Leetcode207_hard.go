package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	temp := make([][]int, numCourses)
	completed := make([]bool, numCourses)
	completing := make([]bool, numCourses)
	hasCycle := false
	for i := 0; i < len(prerequisites); i++ {
		temp[prerequisites[i][0]] = append(temp[prerequisites[i][0]], prerequisites[i][1])
	}
	for i := 0; i < numCourses; i++ {
		if !completed[i] {
			check(temp, i, completed, completing, &hasCycle)
			if hasCycle {
				break
			}
		}
	}
	return !hasCycle
}

func check(temp [][]int, n int, completed, completing []bool, hasCycle *bool) {
	completing[n] = true
	completed[n] = true
	for _, w := range temp[n] {
		if !completed[w] {
			check(temp, w, completed, completing, hasCycle)
		} else if completing[w] {
			*hasCycle = true
			break
		}
	}
	completing[n] = false
}
func main() {
}
