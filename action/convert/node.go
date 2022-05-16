package convert

import (
	"github.com/xsuners/machine/action"
	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/mo/log"
	"go.uber.org/zap"
)

type convert struct {
	node.Node
	props map[string]any
}

func New(s *spec.Node) node.Node {
	return &convert{
		Node:  action.New(s),
		props: s.Props,
	}
}

func (n *convert) Exec(ctx *context.Context) error {
	log.Infosc(ctx, "convert exec", zap.Any("props", n.props))
	for _, data := range n.props {
		// TODO 取值
		switch m := data.(type) {
		case *spec.Create:
			ctx.In.Create = *m
		case *spec.Update:
			ctx.In.Update = *m
		case *spec.Delete:
			ctx.In.Delete = *m
		}
	}
	return nil
}
