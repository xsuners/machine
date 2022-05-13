package machine

import (
	"github.com/nats-io/nats.go"
	"shepin.live/go/machine/spec"
)

type Handler struct {
	Root Node
}

type Machine struct {
	events map[string]*Event
}

func New(s *spec.Specification) *Machine {
	for _, rpc := range s.Rpcs {
		build(rpc.Root)
	}
	return &Machine{}
}

func build(cfg *spec.Node) Node {
	builder, ok := registry[cfg.Name]
	if !ok {
		panic("node not found")
	}
	node := builder(cfg.Config)
	for _, child := range cfg.Children {
		node.Children(build(child))
	}
	return node
}

func (m *Machine) context(in *nats.Msg) *Context {
	return &Context{
		machine: m,
		memory:  newbb(nil),
	}
}

func (m *Machine) Init() {
	// sub(m.events)
}
