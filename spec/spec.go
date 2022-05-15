package spec

import (
	"github.com/xsuners/machine/spec/data"
	"github.com/xsuners/machine/spec/ui"
)

type Mq struct {
	Root    *Node
	Subject string
	// Message Message
}

type Rpc struct {
	Root *Node
	Path string
	// Message Message
}

type Node struct {
	Type     string
	Name     string
	Props    map[string]any
	Children []*Node
}

type Machine struct {
	Rpcs []*Rpc
	Mq   []*Mq
}

type Spec struct {
	Pages     []*ui.Page
	Databases []*data.Database
	Machine   Machine
}
