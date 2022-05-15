package condition

import (
	"fmt"

	"github.com/xsuners/machine/composite"
	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/spec"
)

type condition struct {
	node.Node
	nodes []node.Node
}

func New(c *spec.Node) node.Node {
	return &condition{
		Node: composite.New(c),
	}
}

func (n *condition) Exec(ctx *context.Context) error {
	fmt.Println("sequence exec")
	for _, node := range n.nodes {
		err := node.Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *condition) Children(nodes ...node.Node) map[string]node.Node {
	n.nodes = append(n.nodes, nodes...)
	return n.Node.Children(nodes...)
}
