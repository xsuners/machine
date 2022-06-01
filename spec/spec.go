package spec

type Mq struct {
	Root    *Node
	Subject string
}

type Rpc struct {
	Root *Node
	Path string
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
	Pages     []*Page
	Databases []*Database
	Machine   Machine
}
