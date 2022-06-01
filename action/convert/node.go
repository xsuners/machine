package convert

import (
	"strings"

	"github.com/xsuners/machine/action"
	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/machine/spec/in"
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
		switch m := data.(type) {
		case *in.Create:
			ctx.In.Create = *m
		case *in.Update:
			ctx.In.Update = *m
		case *in.Delete:
			ctx.In.Delete = *m
		}
	}
	if istats, ok := n.props["statements"]; ok {
		if stats, ok := istats.([]string); ok {
			for _, stat := range stats {
				parts := strings.Split(stat, " ")
				switch parts[0] {
				case "asign":
					data, ok := ctx.Get(parts[2])
					if ok {
						err := ctx.Set(parts[1], data)
						if err != nil {
							return err
						}
					}
				case "append":
					data, ok := ctx.Get(parts[2])
					if ok {
						err := ctx.Set(parts[1], data, parts[0])
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
