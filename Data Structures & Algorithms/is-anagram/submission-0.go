func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false;
    }
    wordoccurance  := make(map[byte]int)
    for i:=0;i<len(s);i++ {
        wordoccurance[s[i]]+= 1
    }

    for j:=0;j<len(t);j++ {
        wordoccurance[t[j]]-= 1
    }

    for _, value := range wordoccurance {
		if value != 0 {
            return false
        }
	}
    return true

}
