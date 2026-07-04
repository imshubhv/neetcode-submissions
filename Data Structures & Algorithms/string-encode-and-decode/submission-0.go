type Solution struct{}

	func (s *Solution) Encode(strs []string) string {
		output:= ""
		for _, value := range strs {
			for _, runeVal := range []rune(value){
				output+= string(runeVal+10)
			}
			output+= " "
		}
		return output
	}

func (s *Solution) Decode(encoded string) []string {
	var answer []string 
	tmp:= ""
	for _,value := range []rune(encoded) {
		if value == 32 {
			answer = append(answer, tmp)
			tmp = ""
		} else {
			tmp = tmp + string(value-10)
		}
	}
	return answer
}
