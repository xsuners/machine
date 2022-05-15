package composite

import (
	"fmt"

	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/spec"
)

type composite struct {
	name  string
	props map[string]any
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

func (n *composite) Exec(ctx *context.Context) error {
	fmt.Println("composite exec")
	return nil
}

func (n *composite) Children(nodes ...node.Node) map[string]node.Node {
	for _, node := range n.nodes {
		n.nodes[node.Name()] = node
	}
	return n.nodes
}
