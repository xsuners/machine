package spec

type Spec struct {
	Machine Machine
}

type Machine struct {
	Rpcs []*Rpc
	Mq   []*Mq
}

type Rpc struct {
	Root *Node
	Path string
}

type Mq struct {
	Root    *Node
	Subject string
}

type Node struct {
	Type     string
	Name     string
	Props    map[string]interface{}
	Children []*Node
}
