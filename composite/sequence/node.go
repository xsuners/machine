package sequence

import (
	"fmt"

	"github.com/xsuners/machine/composite"
	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/spec"
)

type sequence struct {
	node.Node
	nodes []node.Node
}

func New(c *spec.Node) node.Node {
	return &sequence{
		Node: composite.New(c),
	}
}

func (n *sequence) Exec(ctx *context.Context) error {
	fmt.Println("sequence exec")
	for _, node := range n.nodes {
		err := node.Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *sequence) Children(nodes ...node.Node) map[string]node.Node {
	n.nodes = append(n.nodes, nodes...)
	return n.Node.Children(nodes...)
}
