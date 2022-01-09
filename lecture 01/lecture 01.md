# 제 1장

올바른 입력이 들어왔을 때의 경우를 살펴보자.

예를 들어 `3+4-6/(2*3)` 이라는 수식이 입력되면,
그 결과로 `6`이 출력된다. 그 과정을 살펴보면,

현재 입력된 수식은 중위 표기법으로 작성되어있다.

**중위 표기법(infix notation)**은 연산자의 우선순위를 정해두고 그 우선순위대로
해당 연산자의 양측 항을 계산하는 방식이다.

이를 그대로 계산하기에는 연산 순서를 나타낸 순서표가 필요하고 로직이 복잡해지는 어려움이 있다.

반면 **후위 표기법(postfix notation)**은 연산자의 우선순위를 정해두지 않고
연산자가 나타나면 앞서 나온 두 항을 해당 연산자로 계산하고 계속해서 이어가는 방식이다.

예를 들어 `3+4-6/(2*3)` 이라는 수식은 후위표기법으로 나타내면 `3 4 + 6 2 3 * / -`가 된다.

연산 과정을 괄호로 묶어 나타내면 다음과 같다.

```
(3 4 +) 6 2 3 * / -
```
3, 4 이후 `+`가 나왔으므로 3, 4를 더한다.
```
7 6 (2 3 *) / -
```
이후 2, 3이 나오고 `*`가 나왔으므로 2, 3을 곱한다.
```
7 (6 6 /) -
```
이후 `/`가 나왔으므로 그 앞의 두 항 6, 6을 서로 나눈다.
```
(7 1 -)
```
이후 `-`이 나왔으므로 앞의 항 7에서 1을 뺀다.
```
6
```
계산 결과는 6이다.

후위 표기법은 연산 순서를 굳이 정의하지 않아도 하나의 규칙으로 계산할 수 있기 때문에 대부분의 계산기는 후위 표기법으로 작성된 수식을 계산하는데 사용한다.

후위 표기법으로 된 수식의 계산은 스택을 사용하여 쉽게 구현할 수 있다.

다음은 go로 작성한 스택 구현 예시이다.
올바른 수식만 들어온다고 가정했으므로 별도의 예외처리는 하지 않았다.

```go
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
```

입력받은 수식은 **중위 표현법**인데, **후위 표기법으로 바꾸어야** 위의 스택을 사용한 방법을 적용할 수 있다.

중위 표기법을 후위 표기법으로 변환하는 과정을 살펴보자.

다음은 중위 표기법을 후위 표기법으로 변환하는 go 함수이다.
연산자 우선순위에 따라 다르게 저장됨을 확인할 수 있다.
```go
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
```

따라서 중위 표현식을 계산하는 식은 다음과 같다.
```go
func calc(expr string) int {
    return stackCalc(infixToPostfix(expr))
}
```

쨘! 우리는 제대로 작동하는 계산기를 만들어냈다.
단, '올바른 수식이 들어왔을 때'만.

다음에는 '올바르지 않은 수식'이 입력으로 주어졌을 때 이를 어떻게 판별해 낼 것인지 생각해보자.