package tree

type Node struct {
	value    interface{}
	children []Node
}

func NewNode(value interface{}, children []Node) *Node {
	return &Node{
		value:    value,
		children: children,
	}
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) AddChild(value interface{}) {
	nc := Node{
		value:    value,
		children: []Node{},
	}
	n.children = append(n.children, nc)
}
