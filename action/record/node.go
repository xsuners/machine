package record

import (
	"fmt"

	"github.com/xsuners/machine/action"
	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/spec"
)

type record struct {
	node.Node
}

func New(s *spec.Node) node.Node {
	return &record{
		Node: action.New(s),
	}
}

func (n *record) Exec(ctx *context.Context) error {
	fmt.Println("record exec", n.Name())
	return nil
}
