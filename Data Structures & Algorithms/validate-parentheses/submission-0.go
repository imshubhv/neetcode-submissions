func isValid(s string) bool {
	stack:= []rune{}

	mapped:= map[rune]rune{
		')':'(',
		'}':'{',
		']':'[',
	}

	for _, value:= range s {
		if value == '(' ||value == '{' ||value == '['{
			stack = append(stack, value)
		}else {
			if len(stack) == 0 || stack[len(stack)-1] != mapped[value]{
				return false
			} 
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
