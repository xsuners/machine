package composite

import (
	"fmt"

	"shepin.live/go/machine/context"
	"shepin.live/go/machine/node"
	"shepin.live/go/machine/spec"
)

type composite struct {
	name  string
	props map[string]interface{}
	nodes map[string]node.Node
}

func New(c *spec.Node) node.Node {
	return &composite{
		name:  c.Name,
		props: c.Props,
	}
}

func (n *composite) Name() string {
	return n.name
}

func (n *composite) Exec(ctx *context.Context) (node.Result, error) {
	fmt.Println("composite exec")
	return node.Success, nil
}

func (n *composite) Children(nodes ...node.Node) map[string]node.Node {
	for _, node := range n.nodes {
		n.nodes[node.Name()] = node
	}
	return n.nodes
}
