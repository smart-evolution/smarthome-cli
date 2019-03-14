package tree

type Tree struct {
	root Node
}

func NewTree(r Node) *Tree {
	return &Tree{
		root: r,
	}
}

type Criterion func(Node) bool
type Algorithm func(Node, Criterion, *Node)

func (t *Tree) SetRoot(n Node) {
	t.root = n
}

func (t *Tree) Search(c Criterion, a Algorithm) Node {
	var searched Node
	a(t.root, c, &searched)
	return searched
}

func TraverseBFS(n Node, c Criterion, s *Node) {
	for _, sn := range n.children {
		if ok := c(sn); ok {
			*s = sn
			return
		}
	}

	for _, sn := range n.children {
		TraverseBFS(sn, c, s)
	}
}
