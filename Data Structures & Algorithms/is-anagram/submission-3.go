func isAnagram(s string, t string) bool {
    if len(s) != len(t){ return false}
    var maps = make(map[byte]int);

    for i:=0; i<len(s); i++ {
        maps[s[i]] = maps[s[i]]+1;
    }
    for i:=0; i<len(s); i++ {
        maps[t[i]] = maps[t[i]]-1;
    }
    for _,freq := range maps {
        if freq !=0 {
            return false
        }
    }
    return true
}
