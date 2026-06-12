
func evalRPN(tokens []string) int {
    stack := make([]int, len(tokens))
    top := -1

    for _, token := range tokens {
        switch token {
        case "+", "-", "*", "/":
            val2 := stack[top]
            val1 := stack[top-1]
            top -= 2

            var result int
            switch token {
            case "+":
                result = val1 + val2
            case "-":
                result = val1 - val2
            case "*":
                result = val1 * val2
            case "/":
                result = val1 / val2
            }

            top++
            stack[top] = result

        default:
            num, _ := strconv.Atoi(token)
            top++
            stack[top] = num
        }
    }

    return stack[top]
}
