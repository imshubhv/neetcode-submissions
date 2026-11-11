type TrieNode struct {
	children [26] *TrieNode
	isEnd	bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{
		root: &TrieNode{},
	}
}

func (this *WordDictionary) AddWord(word string)  {
    current:= this.root

	for _, c := range word{
		index := c-'a'

		if current.children[index] == nil {
			current.children[index] = & TrieNode{}
		}
		current = current.children[index]
	}
	current.isEnd = true
}

func SearchWord(word string, node *TrieNode) bool{
	current:= node

	for index, c := range word{
		if c == '.'{
			for _,ptr := range current.children{
				if ptr != nil && SearchWord(word[index+1:], ptr){
					return true
				}
			}
			return false
		} else {
			index := c - 'a'

			if current.children[index] == nil {
				return false
			}
			current = current.children[index]
		}
	}
	return current.isEnd

}
func (this *WordDictionary) Search(word string) bool {
    return SearchWord(word, this.root)
}
