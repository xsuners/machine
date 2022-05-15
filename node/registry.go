package node

import "github.com/xsuners/machine/spec"

type Builder func(*spec.Node) Node

var registry = make(map[string]Builder)

func Register(name string, builder Builder) {
	registry[name] = builder
}

// func Fetch(name string) (Builder, bool) {
// 	b, ok := registry[name]
// 	return b, ok
// }
