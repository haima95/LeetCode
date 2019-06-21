package main

import "fmt"

func isMatch2(s string, p string) bool {
	n := len(s)
	m := len(p)
	temp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		temp[i] = make([]bool, m+1)
	}
	temp[0][0] = true
	if m > 0 && p[0] == '*' {
		for i := 1; i < n+1; i++ {
			temp[i][0] = true
		}
		for j := 0; j < m && p[j] == '*'; j++ {
			temp[0][j+1] = true
		}
	}

	for j := 1; j <= m; j++ {
		for i := 1; i <= n; i++ {
			if (p[j-1] == s[i-1] || p[j-1] == '?') && temp[i-1][j-1] {
				temp[i][j] = true
			}
			if p[j-1] == '*' {
				if temp[i][j-1] {
					temp[i][j] = true
				}
				if temp[i-1][j] {
					temp[i][j] = true
				}
			}
		}
	}

	fmt.Print("  ")
	for i := 0; i < m; i++ {
		fmt.Printf("%c ", p[i])
	}
	fmt.Println()
	for i := 1; i < n+1; i++ {
		fmt.Printf("%c ", s[i-1])
		for j := 1; j < m+1; j++ {
			if temp[i][j] {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}

		}
		fmt.Println()
	}
	return temp[n][m]
}

func isMatch3(s, p string) bool {
	m := len(s)
	n := len(p)
	i, j := 0, 0
	matchI, startJ := -1, -1
	for i < m {
		if j < n && p[j] == '*' {
			matchI = i
			startJ = j
			j++
		} else if j < n && (p[j] == '?' || p[j] == s[i]) {
			j++
			i++
		} else if startJ > -1 {
			i = matchI + 1
			j = startJ + 1
			matchI++
		} else {
			return false
		}
	}
	for j < n && p[j] == '*' {
		j++
	}
	return j == n
}

type Test struct {
	s string
	p string
	r bool
}

func main() {

	test := []Test{
		{"abefcdgiescdfimde", "ab*cd?i*de", true},
		{"acdcb", "a*c?b", false},
		{"mississippi", "m??*ss*?i*pi", false},
		{"a", "a*", true},
		{"a", "aa", false},
		{"b", "?*?", false},
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
		{"adceb", "*a*b", true},
		{"adceb", "a***c***b", true},
	}

	for _, v := range test {
		if isMatch2(v.s, v.p) == v.r {
			fmt.Printf("s:%s , p:%s , r:%d\n", v.s, v.p, 1)
		} else {
			fmt.Printf("s:%s , p:%s , r:%d\n", v.s, v.p, 0)
		}

	}
}

//  a b * c d ? i * d e
//a 1 0 0 0 0 0 0 0 0 0
//b 0 1 0 0 0 0 0 0 0 0
//e 0 0 1 0 0 0 0 0 0 0
//f 0 0 1 0 0 0 0 0 0 0
//c 0 0 1 1 0 0 0 0 0 0
//d 0 0 1 0 1 0 0 0 0 0
//g 0 0 1 0 0 1 0 0 0 0
//i 0 0 1 0 0 0 1 0 0 0
//e 0 0 1 0 0 0 0 1 0 0
//s 0 0 1 0 0 0 0 1 0 0
//c 0 0 1 1 0 0 0 1 0 0
//d 0 0 1 0 1 0 0 1 1 0
//f 0 0 1 0 0 1 0 1 0 0
//i 0 0 1 0 0 0 1 1 0 0
//m 0 0 1 0 0 0 0 1 0 0
//d 0 0 1 0 0 0 0 1 1 0
//e 0 0 1 0 0 0 0 1 0 1
//s:abefcdgiescdfimde , p:ab*cd?i*de , r:1
//  a * c ? b
//a 1 0 0 0 0
//c 0 1 1 0 0
//d 0 1 0 1 0
//c 0 1 1 0 0
//b 0 1 0 1 0
//s:acdcb , p:a*c?b , r:1
//m ? ? * s s * ? i * p i
//m 1 0 0 0 0 0 0 0 0 0 0 0
//i 0 1 0 0 0 0 0 0 0 0 0 0
//s 0 0 1 0 0 0 0 0 0 0 0 0
//s 0 0 0 1 1 0 0 0 0 0 0 0
//i 0 0 0 1 0 0 0 0 0 0 0 0
//s 0 0 0 1 1 0 0 0 0 0 0 0
//s 0 0 0 1 1 1 0 0 0 0 0 0
//i 0 0 0 1 0 0 1 1 0 0 0 0
//p 0 0 0 1 0 0 1 1 0 0 0 0
//p 0 0 0 1 0 0 1 1 0 0 0 0
//i 0 0 0 1 0 0 1 1 1 1 0 0
//s:mississippi , p:m??*ss*?i*pi , r:1
//a *
//a 1 1
//s:a , p:a* , r:1
//a a
//a 1 0
//s:a , p:aa , r:1
//? * ?
//b 1 1 1
