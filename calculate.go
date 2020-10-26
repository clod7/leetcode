package main

import "fmt"

/*
1. 遇到数字运算
2. 遇到空格跳过
3. +号直接入栈
4. -号负数入栈
5. *号与pop()得到的数乘上一个数
6. /号与pop()得到的数除上一个数
7. 遍历加上栈中的所有数得到结果
*/

func calculate(s string) int {
	var stack []int
	// op,num = 上一个运算符，上一个数字
	var op byte = '+'
	var num int

	s += "p" // 占位符，为了处理最后一个数字
	length := len(s)
	for i := 0; i < length; i++ {
		v, ok := byte2int(s[i])
		if ok {
			num = num*10 + v
		} else if s[i] == ' ' {
			continue
		} else {
			switch op {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				prev := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				stack = append(stack, prev*num)
			case '/':
				prev := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				stack = append(stack, prev/num)
			}

			num = 0
			op = s[i]
		}
	}

	var res int
	for _, num := range stack {
		res += num
	}

	return res
}

func byte2int(b byte) (int, bool) {
	if '0' > b || b > '9' {
		return 0, false
	}

	return int(b - '0'), true
}

func main() {
	fmt.Println(calculate(" 3+5 / 2 "))
}
