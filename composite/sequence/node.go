package sequence

import (
	"fmt"

	"shepin.live/go/machine/composite"
	"shepin.live/go/machine/context"
	"shepin.live/go/machine/node"
	"shepin.live/go/machine/spec"
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

func (n *sequence) Exec(ctx *context.Context) (node.Result, error) {
	fmt.Println("sequence exec")
	for _, node := range n.nodes {
		ret, err := node.Exec(ctx)
		if err != nil {
			return ret, err
		}
	}
	return node.Success, nil
}

func (n *sequence) Children(nodes ...node.Node) map[string]node.Node {
	n.nodes = append(n.nodes, nodes...)
	return n.Node.Children(nodes...)
}
