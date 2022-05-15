package node

import (
	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/spec"
)

type Node interface {
	Name() string
	Exec(ctx *context.Context) error
	Children(...Node) map[string]Node
}

func Build(c *spec.Node) Node {
	builder, ok := registry[c.Type]
	if !ok {
		panic("node not found")
	}
	node := builder(c)
	for _, child := range c.Children {
		node.Children(Build(child))
	}
	return node
}
