package mq

import (
	"shepin.live/go/machine/context"
	"shepin.live/go/machine/node"
	"shepin.live/go/machine/spec"
)

type Handler interface {
	Handle(ctx *context.Context)
	Subject() string
}

func Build(s *spec.Mq) Handler {
	mq := &mq{
		subject: s.Subject,
		root:    node.Build(s.Root),
	}
	return mq
}

type mq struct {
	subject string
	root    node.Node
}

func (m *mq) Subject() string {
	return m.subject
}

func (m *mq) Handle(ctx *context.Context) {
	m.root.Exec(ctx)
}
