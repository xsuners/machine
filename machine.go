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
	"github.com/xsuners/machine/spec/in"
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
	ctx.In = in.In{
		List: in.List{
			Database: "machine",
			Table:    "user",
			Selects: []*in.Select{
				{
					Prop: "id",
					Kind: spec.Int,
				},
				{
					Prop: "name",
					Kind: spec.String,
				},
				{
					Prop: "phone",
					Kind: spec.String,
				},
			},
			Queries: []*in.Query{
				{
					Type:   spec.Gt,
					Prop:   "name",
					Kind:   spec.String,
					Values: []any{"machine"},
				},
				{
					Type:   spec.Lt,
					Prop:   "id",
					Kind:   spec.Int,
					Values: []any{7},
				},
			},
			Page: 0,
			Size: 10,
		},
		Create: in.Create{
			Objects: []*in.Object{
				{
					Database: "machine",
					Table:    "user",
					Props: []*in.Prop{
						{
							Name:  "name",
							Kind:  spec.String,
							Value: "hello",
						},
						{
							Name:  "phone",
							Kind:  spec.String,
							Value: "1313699",
						},
					},
				},
			},
		},
		Update: in.Update{
			Database: "machine",
			Table:    "user",
			Queries: []*in.Query{
				{
					Type:   spec.Eq,
					Prop:   "id",
					Kind:   spec.Int,
					Values: []any{1},
				},
			},
			Props: []*in.Prop{
				{
					Name:  "phone",
					Kind:  spec.String,
					Value: "1200",
				},
			},
		},
		Event: in.Event{
			Props: []*in.Prop{
				{
					Name:  "id",
					Value: 100000,
				},
			},
		},
	}

	for _, r := range m.rpcs {
		r.Handle(ctx)
	}
	for _, m := range m.mqs {
		m.Handle(ctx)
	}
}
