package menu

import (
	"github.com/smart-evolution/smarthome-cli/commands/connect"
	"github.com/smart-evolution/smarthome-cli/commands/status"
	"github.com/smart-evolution/smarthome-cli/tree"
	"net"
)

type Handler func(net.Conn)

type Item struct {
	label   string
	handler Handler
}

func (i *Item) Handler(conn net.Conn) {
	i.handler(conn)
}

func (i *Item) Label() string {
	return i.label
}

type Menu struct {
	tree.Tree
}

func New() *Menu {
	connectItem := Item{
		label:   "connect",
		handler: connect.Handler,
	}
	connectChildren := []tree.Node{}
	connectNode := *tree.NewNode(connectItem, connectChildren)

	statusItem := Item{
		label:   "status",
		handler: status.Handler,
	}
	statusChildren := []tree.Node{}
	statusNode := *tree.NewNode(statusItem, statusChildren)

	rootItem := Item{
		label:   "smarthome",
		handler: func(conn net.Conn) {},
	}
	rootChildren := []tree.Node{
		connectNode,
		statusNode,
	}
	rootNode := *tree.NewNode(rootItem, rootChildren)

	tree := tree.NewTree(rootNode)

	return &Menu{
		*tree,
	}
}

type Commander struct {
	state	[]string
}

func (m *Menu) Execute(cmd []string, conn net.Conn) {
	var cmdHandler Handler

	for _, c := range cmd {
		node := m.Root()

		for _, d := range node.Children() {
			item := d.Value()
			if i, ok := item.(Item); ok {
				if i.Label() == c {

				}
			}
		}
	}

	/*criterion := func(node tree.Node) bool {
		v := node.Value()
		if mi, ok := v.(Item); ok {
			return mi.Label() == cmd
		}
		return false
	}

	cmdNode := m.Search(criterion, tree.TraverseBFS)
	cmdItem := cmdNode.Value()

	if i, ok := cmdItem.(Item); ok {
		i.Handler(conn)
	}*/
}
