package main

import "fmt"

func countAndSay(n int) string {
	temp := map[int]string{
		1: "1", 2: "11", 3: "21", 4: "1211", 5: "111221",
	}
	if n < 6 {
		return temp[n]
	}
	t := []byte(temp[5])
	var result []byte
	for i := 6; i <= n; i++ {
		result = []byte{}
		c := 1
		j := 1
		for ; j < len(t); j++ {
			if t[j] == t[j-1] {
				c++
			} else {
				result = append(result, byte(c+'0'), t[j-1])
				c = 1
			}
		}
		result = append(result, byte(c+'0'), t[j-1])
		t = result
	}
	return string(result)
}

func main() {
	for n := 6; n < 10; n++ {
		fmt.Println(countAndSay(n))
	}
}
