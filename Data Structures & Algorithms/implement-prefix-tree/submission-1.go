type TrieNode struct{
	children [26] *TrieNode
	isEnd	bool
}

type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
    return PrefixTree{
		root: &TrieNode{},
	}
}

func (this *PrefixTree) Insert(word string) {
	current := this.root	

	for _, c := range word {
		index := c - 'a'

		if current.children[index] == nil{
			current.children[index] = &TrieNode{}
		}
		current = current.children[index]
	}
	current.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	current := this.root	

	for _, c := range word {
		index := c - 'a'

		if current.children[index] == nil{
			return false
		}
		current = current.children[index]
	}
	return current.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
current := this.root	

	for _, c := range prefix {
		index := c - 'a'

		if current.children[index] == nil{
			return false
		}
		current = current.children[index]
	}
	return true
}
