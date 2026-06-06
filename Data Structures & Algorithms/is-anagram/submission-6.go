import "slices"
func isAnagram(s string, t string) bool {
    sr:= []rune(s)
    tr:= []rune(t)

    slices.Sort(sr)
    slices.Sort(tr)

    if string(sr) == string(tr){
        return true
    }
    return false
}
