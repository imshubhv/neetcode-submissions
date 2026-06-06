import "slices"
func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    
    sr, tr:= []rune(s), []rune(t)

    slices.Sort(sr)
    slices.Sort(tr)

    if string(sr) == string(tr){
        return true
    }
    return false
}
