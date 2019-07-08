package main

import "fmt"

func fullJustify(words []string, maxWidth int) []string {
	spaces := make([]byte, maxWidth)
	for i := 0; i < maxWidth; i++ {
		spaces[i] = ' '
	}
	var result []string
	n := len(words)
	for i := 0; i < n; i++ {
		if i == n-1 { // 如果是最后一个直接在后面补充' '，并存入数组
			result = append(result, words[i]+string(spaces[:maxWidth-len(words[i])]))
		} else {
			temp := len(words[i])
			tempStr := []string{words[i]}
			j := i + 1
			for ; j < n; j++ {
				if temp+len(tempStr)+len(words[j]) > maxWidth { // 先假设字符串间以一个空格隔开，如果加上下一个字符串会大于maxWidth，则不继续加了
					if len(tempStr) == 1 { // 如果只有一个字符串，其实就是一个长度很大的字符串，则直接在后面补充' '，并存入数组
						result = append(result, tempStr[0]+string(spaces[:maxWidth-len(tempStr[0])]))
					} else {
						step := (maxWidth - temp) / (len(tempStr) - 1) // 因为要尽量均匀分别空格，则先计算字符串间平均可以放几个
						temp = (maxWidth - temp) % (len(tempStr) - 1)  // 求出剩下的空格数
						th := ""
						for k := 0; k < len(tempStr)-1; k++ {
							if k < temp {
								th += tempStr[k] + string(spaces[:step+1]) // 因为左侧空格数要大于右侧，又因为空格尽量要均匀分配，所以剩下的空格在左侧的字符串间再多加一个空格
							} else {
								th += tempStr[k] + string(spaces[:step])
							}
						}
						th += tempStr[len(tempStr)-1]
						result = append(result, th)
					}
					i = j - 1
					break
				} else {
					temp += len(words[j])
					tempStr = append(tempStr, words[j])
					if j == n-1 { // 如果是最后一个
						th := ""
						for k := 0; k < len(tempStr)-1; k++ {
							th += tempStr[k] + string(spaces[:1])
						}
						th += tempStr[len(tempStr)-1]
						if maxWidth-temp-len(tempStr) > -1 { // 加了空格后字符个数有没有达到maxWith，没有则补上
							th += string(spaces[:maxWidth-temp-len(tempStr)+1])
						}
						result = append(result, th)
						i = j
						break
					}
				}
			}
		}
	}
	return result
}

func main() {
	//words := []string{"This", "is", "ok"}
	//words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	//words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	//words := []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"}
	words := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	result := fullJustify(words, 16)
	for i := 0; i < len(result); i++ {
		fmt.Printf("[%s]", result[i])
	}
}
