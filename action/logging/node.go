package logging

import (
	"github.com/xsuners/machine/action"
	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/mo/log"
	"go.uber.org/zap"
)

type logging struct {
	node.Node
	props map[string]any
}

func New(s *spec.Node) node.Node {
	return &logging{
		Node:  action.New(s),
		props: s.Props,
	}
}

func (n *logging) Exec(ctx *context.Context) error {
	log.Infosc(ctx, "logging exec", zap.Any("props", n.props))

	for k, v := range n.props {
		val, ok := ctx.Get(v.(string))
		if ok {
			log.Infosc(ctx, "logging exec", zap.Any(k, val))
		} else {
			log.Warnsc(ctx, "logging exec", zap.Any(k, val))
		}
	}

	return nil
}
