package main

import "fmt"

func recIsValid(s string, stack []byte) bool {
	if len(s) == 0 {
		return len(stack) == 0
	}
	if s[0] == ')' || s[0] == ']' || s[0] == '}' {
		if len(stack) == 0 {
			return false
		}
		lstInStack := stack[len(stack)-1]
		if (lstInStack == '(' && s[0] == ')') ||
			(lstInStack == '[' && s[0] == ']') ||
			(lstInStack == '{' && s[0] == '}') {
			stack = stack[0 : len(stack)-1]
			return recIsValid(s[1:], stack)
		} else {
			return false
		}
	} else {
		stack = append(stack, s[0])
		return recIsValid(s[1:], stack)
	}
}

func isValid(s string) bool {
	stack := make([]byte, 0)
	return recIsValid(s, stack)
}

func main() {
	fmt.Println(isValid("{}[](([]()[]))"))
}
