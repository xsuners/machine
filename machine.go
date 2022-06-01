package machine

import (
	"fmt"

	"github.com/xsuners/machine/action/convert"
	"github.com/xsuners/machine/action/create"
	"github.com/xsuners/machine/action/list"
	"github.com/xsuners/machine/action/logging"
	"github.com/xsuners/machine/action/record"
	"github.com/xsuners/machine/action/update"
	"github.com/xsuners/machine/composite/sequence"
	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/node"
	"github.com/xsuners/machine/server/mq"
	"github.com/xsuners/machine/server/rpc"
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/machine/spec/types"
)

type Machine struct {
	c    *spec.Machine
	mqs  []mq.Handler
	rpcs []rpc.Handler
}

func New(s *spec.Machine) *Machine {
	m := &Machine{
		c: s,
	}
	for _, c := range s.Mq {
		m.mqs = append(m.mqs, mq.Build(c))
	}
	for _, c := range s.Rpcs {
		m.rpcs = append(m.rpcs, rpc.Build(c))
	}
	return m
}

func init() {
	// action
	node.Register("list", list.New)
	node.Register("record", record.New)
	node.Register("create", create.New)
	node.Register("update", update.New)
	node.Register("convert", convert.New)
	node.Register("logging", logging.New)
	// composite
	node.Register("sequence", sequence.New)
}

func (m *Machine) Boot() {
	fmt.Println("boot")
	ctx := context.New()
	ctx.In = spec.In{
		List: spec.List{
			Database: "machine",
			Table:    "user",
			Selects: []*spec.Select{
				{
					Prop: "id",
					Kind: types.Int,
				},
				{
					Prop: "name",
					Kind: types.String,
				},
				{
					Prop: "phone",
					Kind: types.String,
				},
			},
			Queries: []*spec.Query{
				{
					Type:   types.Gt,
					Prop:   "name",
					Kind:   types.String,
					Values: []any{"machine"},
				},
				{
					Type:   types.Lt,
					Prop:   "id",
					Kind:   types.Int,
					Values: []any{7},
				},
			},
			Page: 0,
			Size: 10,
		},
		Create: spec.Create{
			Objects: []*spec.Object{
				{
					Database: "machine",
					Table:    "user",
					Props: []*spec.Prop{
						{
							Name:  "name",
							Kind:  types.String,
							Value: "hello",
						},
						{
							Name:  "phone",
							Kind:  types.String,
							Value: "1313699",
						},
					},
				},
			},
		},
		Update: spec.Update{
			Database: "machine",
			Table:    "user",
			Queries: []*spec.Query{
				{
					Type:   types.Eq,
					Prop:   "id",
					Kind:   types.Int,
					Values: []any{1},
				},
			},
			Props: []*spec.Prop{
				{
					Name:  "phone",
					Kind:  types.String,
					Value: "1200",
				},
			},
		},
		Event: spec.Event{
			Props: []*spec.Prop{
				{
					Name:  "id",
					Value: 100000,
				},
			},
		},
	}
	// for i := 0; i < 2; i++ {
	for _, m := range m.mqs {
		m.Handle(ctx)
	}
	// for _, r := range m.rpcs {
	// 	r.Handle(ctx)
	// }
	// 	time.Sleep(time.Second * 5)
	// }
}
