package tree

type Tree struct {
    root    Node
}

func NewTree(r Node) *Tree {
    return &Tree {
        root: r,
    }
}

type Criterion func(Node) bool

func (t *Tree) SetRoot(n Node) {
    t.root = n
}

func (t *Tree) Search(c Criterion) Node {
    var searched Node
    traverseBFS(t.root, c, &searched)
    return searched
}

func traverseBFS(n Node, c Criterion, s *Node) {
    for _, sn := range n.children {
        if ok := c(sn); ok {
            *s = sn
            break;
        }
    }

    for _, sn := range n.children {
        traverseBFS(sn, c, s)
    }
}
