package menu

import (
    "github.com/smart-evolution/smarthome-cli/menu/status"
    "github.com/smart-evolution/smarthome-cli/tree"
    "net"
)

type Handler func(net.Conn)

type Item struct {
    label       string
    handler     Handler
}

func (i *Item) Handler(conn net.Conn) {
    i.handler(conn)
}

func (i *Item) Label() string {
    return i.label
}

type Menu struct{
    tree.Tree
}

func New() *Menu {
    statusItem := Item{
        label: "status",
        handler: status.Handler,
    }
    statusChildren := []tree.Node{}
    statusNode := *tree.New(statusItem, statusChildren)

    rootItem := Item {
        label: "smarthome",
        handler: func(conn net.Conn) {},
    }
    rootChildren := []tree.Node{
        statusNode,
    }
    rootNode := *tree.New(rootItem, rootChildren)

    tree := tree.NewTree(rootNode)

    return &Menu{
        *tree,
    }
}

func (m *Menu) Execute(cmd string, conn net.Conn) {
    criterion := func(node tree.Node) bool {
        v := node.Value()
        if mi, ok := v.(Item); ok {
            return mi.Label() == cmd
        }
        return false
    }

    cmdNode := m.Search(criterion)
    cmdItem := cmdNode.Value()

    if i, ok := cmdItem.(Item); ok {
        i.Handler(conn)
    }
}

