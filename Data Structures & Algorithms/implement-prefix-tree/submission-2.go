type node struct {
	children map[rune]*node
	isEndWord bool
}

func newNode() *node {
	return &node{
		children: map[rune]*node{},
	}
}

func (n *node) addChild(c rune, isEndWord bool) *node {
	if _, ok := n.children[c]; !ok {
		n.children[c] = newNode()
	}
	if !n.children[c].isEndWord {
		n.children[c].isEndWord = isEndWord
	}
	return n.children[c]
}

type PrefixTree struct {
	node *node
}

func Constructor() PrefixTree {
    return PrefixTree{
		node: newNode(),
	}
}

func (this *PrefixTree) Insert(word string) {
	l := len(word)
	n := this.node
	for i, c := range word {
		var isEndWord bool
		if i == l-1	{
			isEndWord = true
		}
		n = n.addChild(c, isEndWord)
	}
}

func (this *PrefixTree) Search(word string) bool {
	var ok bool
	n := this.node
	for _, c := range word {
		n, ok = n.children[c]
		if !ok {
			return false
		}
	}
	return n.isEndWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	var ok bool
	n := this.node
	for _, c := range prefix {
		n, ok = n.children[c]
		if !ok {
			return false
		}
	}
	return true
}
