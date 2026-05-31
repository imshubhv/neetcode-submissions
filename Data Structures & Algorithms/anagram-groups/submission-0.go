import "slices"
func groupAnagrams(strs []string) [][]string {

    anagramGroup:= make(map[string][]string)

    for _,str:= range strs{
        runes:= []rune(str)
        slices.Sort(runes)
        sorted:= string(runes)
        anagramGroup[sorted] = append(anagramGroup[sorted],str)
    }

    answer:= make([][]string, 0, len(anagramGroup))
    for _, groups := range anagramGroup {
        answer = append(answer, groups)
    }

    return answer
}
