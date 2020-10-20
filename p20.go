package main

import "log"

func isValid(s string) bool {

	map_ := map[byte]byte {
		'}': '{',
		')': '(',
		']': '[',
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == map_[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	log.Println(isValid("{}"))
}
