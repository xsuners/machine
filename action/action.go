package action

import (
	"fmt"

	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/spec"
)

type action struct {
	name  string
	props map[string]any
	nodes map[string]node.Node
}

func New(c *spec.Node) node.Node {
	return &action{
		name:  c.Name,
		props: c.Props,
		nodes: make(map[string]node.Node),
	}
}

func (n *action) Name() string {
	return n.name
}

func (n *action) Exec(ctx *context.Context) error {
	fmt.Println("action exec")
	return nil
}

func (n *action) Children(nodes ...node.Node) map[string]node.Node {
	return n.nodes
}
