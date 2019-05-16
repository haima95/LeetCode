package main

import "fmt"

func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}
	digMap := map[byte]string{
		'2':"abc",
		'3':"def",
		'4':"ghi",
		'5':"jkl",
		'6':"mno",
		'7':"pqrs",
		'8':"tuv",
		'9':"wxyz",
	}
	count :=1
	lenx := 0
	temp := [][]byte{}
	for _,v := range []byte(digits) {
		if v == '1' {
			continue
		}
		count *= len(digMap[v])
		lenx ++
	}
	for i:=0;i<count;i++ {
		temp = append(temp,[]byte{})
	}
	for _,v := range []byte(digits) {
		if v == '1' {
			continue
		}
		str := digMap[v]
		t := 0
		for t < len(temp) {
			for _,k := range []byte(str) {
				for i:=0;i<count/len(str);i++ {
					temp[t] = append(temp[t],k)
					t ++
				}
			}
		}
		count /= len(str)
	}
	for i:=0;i<len(temp);i++ {
		result = append(result,fmt.Sprintf("%s",temp[i]))
	}
	return result
}

func main() {
	//["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
	s := ""
	fmt.Println(letterCombinations(s))
}