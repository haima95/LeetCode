package main

import "fmt"

func lengthOfLongestSubstring(s string) int {  //myfunction
	if len(s) < 2 {
		return len(s)
	}
	result := 0
	n := 0
	for i:=1;i<len(s);i++ {
		for j:= n;j<i;j++ {
			if s[i] == s[j] {
				if i-n > result {
					result = i-n
				}
				n ++
				i--
			}
		}
	}
	fmt.Println(result)
	if len(s)-n > result {
		result = len(s)-n
	}
	return result
}
/**
1 lastOccurred[x]不存在，或者 <start  无需操作
2 lastOccurred[x] >= start 更新start
3 更新lastOccurred[x]和maxLength
 */

func lengthOfLongestSubstring2(s string) int {  // other function
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i,v := range []byte(s) {
		if lastI,ok := lastOccurred[v];ok && lastI >= start {
			start = lastI+1
		}
		if i - start+1 > maxLength {
			maxLength = i - start+1
		}
		lastOccurred[v] = i
	}
	return maxLength
}

func main(){
	s := "pwwabcdbbb"
	fmt.Println(lengthOfLongestSubstring2(s))
}
