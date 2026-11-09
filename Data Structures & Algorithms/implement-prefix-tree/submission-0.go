type Node struct {
	children [26]*Node
	isEnd bool
}
type PrefixTree struct {
	root *Node
}

func Constructor() PrefixTree {
    return PrefixTree{
		root: &Node{},
	}
}

func (this *PrefixTree) Insert(word string) {
	current:= this.root

	for _,ch := range word {
		index:= ch-'a'

		if current.children[index] == nil{
			current.children[index] = &Node{}
		}
		current = current.children[index]
	}
	current.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	current:= this.root

	for _,ch := range word {
		index:= ch-'a'

		if current.children[index] == nil{
			return false
		}
		current = current.children[index]
	}
	return current.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	current:= this.root

	for _,ch := range prefix {
		index:= ch-'a'

		if current.children[index] == nil{
			return false
		}
		current = current.children[index]
	}
	return true
}
