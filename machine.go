package machine

import (
	"fmt"
	"time"

	"shepin.live/go/machine/action/record"
	"shepin.live/go/machine/composite/sequence"
	"shepin.live/go/machine/context"
	"shepin.live/go/machine/node"
	"shepin.live/go/machine/server/mq"
	"shepin.live/go/machine/server/rpc"
	"shepin.live/go/machine/spec"
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
	node.Register("sequence", sequence.New)
	node.Register("record", record.New)
}

func (m *Machine) Boot() {
	fmt.Println("boot")
	for i := 0; i < 10; i++ {
		for _, m := range m.mqs {
			m.Handle(context.New())
		}
		for _, r := range m.rpcs {
			r.Handle(context.New())
		}
		time.Sleep(time.Second * 5)
	}
}
