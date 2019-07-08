package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	var temp []byte
	na := len(a)
	nb := len(b)
	if na < nb {
		na, nb = nb, na
		a, b = b, a
	}
	temp = make([]byte, na+1)
	i := 1
	var t uint8 = 0
	for i <= na {
		if i <= nb {
			t = t + (a[na-i] - '0') + (b[nb-i] - '0')
		} else if i <= na && i > nb {
			t = t + (a[na-i] - '0')

		}
		temp[na-i+1] = (t % 2) + '0'
		t = t / 2
		i++
	}
	if t == 1 {
		temp[0] = '1'
		return string(temp)
	} else {
		return string(temp[1:])
	}
}

type AddBinary struct {
	a string
	b string
	c string
}

func main() {
	t := []AddBinary{
		{"0", "0", "0"},
		{"0", "1", "1"},
		{"1", "1", "10"},
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}
	for i := 0; i < len(t); i++ {
		fmt.Println(addBinary(t[i].a, t[i].b) == t[i].c)
	}

}
