package main

func infix2postfix(infix string) string {
	var postfix string
	var stack []rune
	for _, c := range infix {
		if c == '(' {
			stack = append(stack, c)
		} else if c == ')' {
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				postfix += string(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		} else if c == '+' || c == '-' || c == '*' || c == '/' {
			for len(stack) > 0 &&
				stack[len(stack)-1] != '(' &&
				(c == '*' || c == '/') &&
				(stack[len(stack)-1] == '*' ||
					stack[len(stack)-1] == '/') {
				postfix += string(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, c)
		} else {
			postfix += string(c)
		}
	}
	for len(stack) > 0 {
		postfix += string(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return postfix
}

func stackCalc(expr string) int {
	var stack []int
	for _, c := range expr {
		if c >= '0' && c <= '9' {
			stack = append(stack, int(c-'0'))
		} else {
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch c {
			case '+':
				stack = append(stack, b+a)
			case '-':
				stack = append(stack, b-a)
			case '*':
				stack = append(stack, b*a)
			case '/':
				stack = append(stack, b/a)
			}
		}
	}
	return stack[0]
}

func calc(expr string) int {
	return stackCalc(infix2postfix(expr))
}

func main() {
	println(calc("1+2*3"))
}
