func isPalindrome(s string) bool {
	array:= []string{}
	for _, val:= range s {
		if isAlphaNum(val) {
			if val<=90{
			 val = val+32
			}
			array = append(array, string(val))
		}
	}
	size:= len(array)-1;
	for i:= range array {
		j:= size-i
		if i == j {
			break
		} 

		if array[i] != array[j] {
			return false
		}
	}
	return true
}


func isAlphaNum(c rune) bool {
    return unicode.IsLetter(c) || unicode.IsDigit(c)
}
