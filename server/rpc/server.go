package rpc

import (
	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/spec"
)

type Handler interface {
	Handle(ctx *context.Context)
	Path() string
}

func Build(s *spec.Rpc) Handler {
	r := &rpc{
		root: node.Build(s.Root),
		path: s.Path,
	}
	return r
}

type rpc struct {
	path string
	root node.Node
}

func (r *rpc) Path() string {
	return r.path
}

func (r *rpc) Handle(ctx *context.Context) {
	r.root.Exec(ctx)
}
