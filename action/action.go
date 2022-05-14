package action

import (
	"fmt"

	"shepin.live/go/machine/context"
	"shepin.live/go/machine/node"
	"shepin.live/go/machine/spec"
)

type action struct {
	name  string
	props map[string]interface{}
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

func (n *action) Exec(ctx *context.Context) (node.Result, error) {
	fmt.Println("action exec")
	return node.Success, nil
}

func (n *action) Children(nodes ...node.Node) map[string]node.Node {
	return n.nodes
}
