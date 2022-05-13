package main

import (
	"shepin.live/go/machine"
	"shepin.live/go/machine/spec"
)

func main() {
	c := &spec.Specification{}
	m := machine.New(c)
	m.Init()
}
