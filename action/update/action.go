package update

import (
	"github.com/xsuners/machine/action"
	"github.com/xsuners/machine/action/util"
	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/database"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/mo/log"
	"go.uber.org/zap"
)

type update struct {
	node.Node
}

func New(s *spec.Node) node.Node {
	return &update{
		Node: action.New(s),
	}
}

func (n *update) Exec(ctx *context.Context) error {
	u := ctx.In.Update
	set, vals := util.Set(u.Props)
	cons, vals := util.Where(u.Queries, vals...)
	q := "update " + u.Table +
		" set " + set +
		" where " + cons
	_, _, err := database.Fetch(u.Database).Exec(ctx, q, vals...)
	if err != nil {
		log.Errorsc(ctx, "create exec", zap.Error(err))
		return err
	}
	return nil
}
