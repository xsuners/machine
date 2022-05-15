package list

import (
	"strings"

	"github.com/xsuners/machine/action"
	"github.com/xsuners/machine/action/util"
	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/database"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/mo/log"
	"go.uber.org/zap"
)

type list struct {
	node.Node
}

func New(s *spec.Node) node.Node {
	return &list{
		Node: action.New(s),
	}
}

func (n *list) Exec(ctx *context.Context) error {
	l := ctx.In.List
	cons, vals := util.Where(l.Queries)
	var selects []string
	for _, s := range l.Selects {
		selects = append(selects, s.Prop)
	}
	q := "select " + strings.Join(selects, ",") +
		" from " + l.Table +
		" where " + cons +
		" limit ?,?"
	vals = append(vals, l.Page*l.Size, l.Size)
	rows, err := database.Fetch(l.Database).Query(ctx, q, vals...)
	if err != nil {
		log.Errorsc(ctx, "list exec", zap.Error(err))
		return err
	}
	defer rows.Close()

	var out []any
	for rows.Next() {
		data, vals := util.Scan(l.Selects...)
		err = rows.Scan(vals...)
		if err != nil {
			log.Errorsc(ctx, "list exec", zap.Error(err))
		} else {
			out = append(out, data)
		}
	}
	// d, _ := json.Marshal(out)
	// fmt.Println(">>>", string(d))
	err = ctx.Set("out", out)
	if err != nil {
		log.Errorsc(ctx, "list exec", zap.Error(err))
		return err
	}
	return nil
}
