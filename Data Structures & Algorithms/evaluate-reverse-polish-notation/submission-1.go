func eval(val1 int, val2 int, oper string) int {
	switch oper{
		case "+" :
			return val1 + val2
		case "-" :
			return val1 - val2
		case "*" :
			return val1 * val2
		case "/" :
			return val1 / val2
	}
	return 0
}

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, value := range tokens {
		if value != "+" && value!= "-" && value != "*" && value!= "/"{
			number, _:= strconv.Atoi(value)
			stack = append(stack, number)
		}else {
			val2:= stack[len(stack)-1]
			val1:= stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			val3:= eval(val1, val2, value)
			stack = append(stack, val3)
		}
	}
	return stack[len(stack)-1]
}


