import "slices"
func groupAnagrams(strs []string) [][]string {

    anagramMap:= make(map[string][]string)

    for _, value := range strs {

        runes:= []rune(value)
        slices.Sort(runes)
        sortedString:= string(runes)
        anagramMap[sortedString] = append(anagramMap[sortedString],value)
    }

    resultArray:= make([][]string,0,len(anagramMap))
    for _, value:= range anagramMap {
        resultArray = append(resultArray, value)
    }
    return resultArray
}
