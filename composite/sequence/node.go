package sequence

import (
	"shepin.live/go/machine"
)

type sequence struct {
	props map[string]interface{}
	nodes []machine.Node
}

func New(props map[string]interface{}) machine.Node {
	return &sequence{
		props: props,
	}
}

func (s *sequence) Id() string

func (s *sequence) Exec(ctx *machine.Context) (machine.Result, error) {
	return "", nil
}

func (s *sequence) Children(nodes ...machine.Node) map[string]machine.Node {
	s.nodes = append(s.nodes, nodes...)
	out := make(map[string]machine.Node)
	for _, node := range s.nodes {
		out[node.Id()] = node
	}
	return out
}
