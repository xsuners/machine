package record

import (
	"fmt"

	"shepin.live/go/machine/action"
	"shepin.live/go/machine/context"
	"shepin.live/go/machine/node"
	"shepin.live/go/machine/spec"
)

type record struct {
	node.Node
}

func New(s *spec.Node) node.Node {
	return &record{
		Node: action.New(s),
	}
}

func (n *record) Exec(ctx *context.Context) (node.Result, error) {
	fmt.Println("record exec", n.Name())
	return node.Success, nil
}
