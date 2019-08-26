package main

import "fmt"

func minCut(s string) int {
	local := make([]int, len(s)+1)
	for i := 0; i < len(local); i++ {
		local[i] = i - 1
	}
	for i := 1; i < len(s); i++ {
		for j := 0; i-j >= 0 && i+j < len(s) && s[i+j] == s[i-j]; j++ {
			if local[i+j+1] > local[i-j]+1 {
				local[i+j+1] = local[i-j] + 1
			}
		}
		for j := 0; i-j-1 >= 0 && i+j < len(s) && s[i+j] == s[i-j-1]; j++ {
			if local[i+j+1] > local[i-j-1]+1 {
				local[i+j+1] = local[i-j-1] + 1
			}
		}
	}
	return local[len(s)]
}

func main() {
	s := ""
	for {
		n, err := fmt.Scanf("%s", &s)
		if n == 0 || err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(minCut(s))
	}

	//s := "apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp"
	//fmt.Println(minCut(s))
}
