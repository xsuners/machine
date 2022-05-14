package node

import (
	"shepin.live/go/machine/context"
	"shepin.live/go/machine/spec"
)

type Result string

const (
	Success Result = "success"
	Failure Result = "failure"
	Running Result = "running"
)

type Node interface {
	Name() string
	Exec(ctx *context.Context) (Result, error)
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
