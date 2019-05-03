package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	start := 0
	end := 0
	length := 0
	for i:=0;i<n;i++ {
		tmps := i
		tmpe := i
		// 奇数回文
		for ;tmps >= 0 && tmpe < n; {
			if s[tmps] == s[tmpe] {
				tmps --
				tmpe ++
			}else {
				if tmpe-tmps-1 > length {
					length = tmpe-tmps-1
					start = tmps+1
					end = tmpe-1
				}
				break
			}
		}
		if !(tmps > 0 && tmpe < n-1) && tmpe-tmps-1 > length {
			length = tmpe-tmps-1
			start = tmps+1
			end = tmpe-1
		}
		tmps = i
		tmpe = i+1
		// 偶数回文
		for ;tmps >=0 && tmpe < n; {
			if s[tmps] == s[tmpe] {
				tmps --
				tmpe ++
			}else {
				if tmpe-tmps-1 > length {
					length = tmpe-tmps-1
					start = tmps+1
					end = tmpe-1
				}
				break
			}
		}
		if !(tmps > -1 && tmpe < n) && tmpe-tmps-1 > length {
			length = tmpe-tmps-1
			start = tmps+1
			end = tmpe-1
		}

	}
	return s[start:end+1]
}

func main(){
	s := "bbbb"
	fmt.Println(longestPalindrome(s))
}
