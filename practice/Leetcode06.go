package main

import "fmt"

func convert(s string, numRows int) string {
	n := len(s)
	temp := make([]uint8,n)
	count := 0
	step := 2*numRows-2
	if step == 0 {
		step = 1
	}
	for i:=0;i<numRows;i++ {
		for j:=i;j<n && count<n;j+= step {
			temp[count] = s[j]
			count ++
			if i > 0 && i < numRows-1 && j+2*(numRows-i-1) < n{
				temp[count] = s[j+2*(numRows-i-1)]
				count ++
			}
		}
	}
	return fmt.Sprintf("%s",temp)
}

func main() {
	s := "ABCD"
	fmt.Println(convert(s,2))
}
