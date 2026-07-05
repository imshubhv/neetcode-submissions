func lengthOfLongestSubstring(s string) int {
	size:= len(s)
	if size < 2 {
		return size
	}
	maxLength:= 0

	for i:=0; i<size; i++ {
		myMap := map[byte]bool{}
		myMap[s[i]]=true
		currentLength:=1
		for j:=i+1; j<size; j++ {
			if myMap[s[j]] {
				if currentLength > maxLength {
					maxLength = currentLength
				}
				break
			}
			currentLength++
			if currentLength > maxLength {
					maxLength = currentLength
				}
			myMap[s[j]]=true
		}
	} 
	return maxLength
}
