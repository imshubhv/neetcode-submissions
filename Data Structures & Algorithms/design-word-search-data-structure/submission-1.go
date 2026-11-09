type TrieNode struct {
	children	[26] *TrieNode
	isEnd		bool
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
    current := this.root

	for _, ch := range word {

		index := ch -'a'

		if current.children[index] == nil {
			current.children[index] = &TrieNode{}
		}
		current = current.children[index]
	}

	current.isEnd = true
}
func (this *WordDictionary) Search(word string) bool {
    return search(word, this.root)
}

func search(word string, node *TrieNode) bool {
    current := node

	for index, ch := range word{
		if ch == '.'{
			for _, node := range current.children{
				if node != nil {
					if search(word[index+1:], node) {
						return true
					}
				} 
			}
			return false 
		} else {
			index := ch - 'a'

			if (current.children[index] == nil){
				return false
			}
			current = current.children[index]
		}
	}
	return current.isEnd
}
