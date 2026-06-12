func isPalindrome(s string) bool {
//32
	array:= []string{}
	for _, val:= range s {
		if isAlphaNum(val) {
			if val<=90{
			 val = val+32
			}
			array = append(array, string(val))
		}
	}
	fmt.Println(array)
	size:= len(array)-1;
	for i:= range array {
		j:= size-i
		fmt.Println("i and j are ", i,j)
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
