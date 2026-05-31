import "slices"
func groupAnagrams(strs []string) [][]string {
    anagramMap:= make(map[string][]string)

    for _, str:= range strs{
        runes:= []rune(str)
        slices.Sort(runes)
        key:= string(runes)
        anagramMap[key] = append( anagramMap[key], str)
    }

    result:= make([][]string, 0, len(anagramMap))

    for _, value:= range anagramMap{
        result = append(result, value)
    }

    return result
}
