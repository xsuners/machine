package spec

type Machine struct {
	States []State
	Api    []Api
}

type Param struct {
	Name       string
	Type       string
	Validation string
}
type Body struct{}

type Api struct {
	Url    string
	Params []Param
	Body   Body
}

type Action struct{}

type State struct {
	Name   string
	Code   string
	Events []Event
}

type Event struct {
	Action []Action
}

type Specification struct {
	Machine Machine

	Rpcs []*Rpc
	Mq   []*Mq
}

type Rpc struct {
	Root *Node
}

type Mq struct {
	Root *Node
}

type Node struct {
	Name     string
	Config   map[string]interface{}
	Children []*Node
}
