package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	temp := strings.Split(path, "/")
	for i := 0; i < len(temp); i++ {
		if temp[i] == "." || temp[i] == "" {
			if i == 0 {
				temp = temp[i+1:]

			} else if i < len(temp)-1 {
				temp = append(temp[:i], temp[i+1:]...)
			} else {
				temp = temp[:i]
			}
			i--
		} else if temp[i] == ".." {
			if i < 2 {
				if len(temp) < 2 {
					temp = []string{}
				} else {
					temp = temp[i+1:]
				}
				i = -1
			} else {
				if i < len(temp)-1 {
					temp = append(temp[:i-1], temp[i+1:]...)
				} else {
					temp = temp[:i-1]
				}
				i -= 2
			}
		}
	}
	return "/" + strings.Join(temp, "/")
}

func main() {
	tests := []struct {
		in, out string
	}{
		{"/home", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
	}
	for _, it := range tests {
		t := simplifyPath(it.in)
		if it.out == t {
			fmt.Printf("ok : in:%s,expect:[%s,%d],out:[%s,%d]\n", it.in, t, len(t), it.out, len(it.out))
		} else {
			fmt.Printf("fail : in:%s,expect:[%s,%d],out:[%s,%d]\n", it.in, t, len(t), it.out, len(it.out))
		}
	}

}
