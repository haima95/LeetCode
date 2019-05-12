package main

import (
	"fmt"
)

func intToRoman(num int) string {
	temp := map[int]byte{
		1:'I',
		2:'V',     // 5
		10:'X',
		11:'L',    // 50
		100:'C',
		101:'D',   // 500
		1000:'M',
	}
	result := []byte{}
	t := 1
	for num > 0 {
		k := num % 10
		fmt.Println(k,num,t)
		if t >= 1000 {
			for i := 0; i < k; i++ {
				result = append(result, temp[1000])
			}
		}else {
			if k < 4 {
				for i := 0; i < k; i++ {
					result = append(result, temp[t])
				}
			} else if k == 4 {
				result = append(result, temp[t+1])
				result = append(result, temp[t])
			} else if k > 4 && k < 9 {
				for i := 0; i < k-5; i++ {
					result = append(result, temp[t])
				}
				result = append(result, temp[t+1])
			} else {
				result = append(result, temp[t*10])
				result = append(result, temp[t])
			}
		}
		t *= 10
		num = (num/10)
	}
	n := len(result)
	for i:=0;i<n/2;i++ {
		result[i],result[n-i-1] = result[n-i-1],result[i]
	}
	r := string(result)
	return r
}

func intToRoman2(num int) string {
	S1 := []string{"","I","II","III","IV","V","VI","VII","VIII","IX"}
	S10 := []string{"","X","XX","XXX","XL","L","LX","LXX","LXXX","XC"}
	S100 := []string{"","C","CC","CCC","CD","D","DC","DCC","DCCC","CM"}
	S1000 := []string{"","M","MM","MMM"}
	return S1000[num/1000]+S100[num%1000/100]+S10[num%100/10]+S1[num%10]
}

func main() {
	i := 58
	fmt.Println(intToRoman2(i))
}
