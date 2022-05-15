package create

import (
	"strings"

	"github.com/xsuners/machine/action"
	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/database"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/mo/log"
	"go.uber.org/zap"
)

type create struct {
	node.Node
}

func New(s *spec.Node) node.Node {
	return &create{
		Node: action.New(s),
	}
}

func (n *create) Exec(ctx *context.Context) error {
	for _, object := range ctx.In.Create.Objects {
		var keys []string
		var sets []string
		var vals []any
		for _, prop := range object.Props {
			keys = append(keys, prop.Name)
			sets = append(sets, "?")
			vals = append(vals, prop.Value)
		}
		q := "insert into " + object.Table + "(" + strings.Join(keys, ", ") + ") values(" + strings.Join(sets, ",") + ")"
		_, _, err := database.Fetch(object.Database).Exec(ctx, q, vals...)
		if err != nil {
			log.Errorsc(ctx, "create exec", zap.Error(err))
			return err
		}
	}
	return nil
}
