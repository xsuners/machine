package main

import (
	"shepin.live/go/machine"
	"shepin.live/go/machine/spec"
)

func main() {
	// s := spec.Load([]byte("{}"))
	s := &spec.Spec{
		Machine: spec.Machine{
			Rpcs: []*spec.Rpc{
				{
					Path: "test",
					Root: &spec.Node{
						Type:  "sequence",
						Name:  "s1",
						Props: make(map[string]interface{}),
						Children: []*spec.Node{
							{
								Type:  "record",
								Name:  "r1",
								Props: make(map[string]interface{}),
							},
							{
								Type:  "record",
								Name:  "r2",
								Props: make(map[string]interface{}),
							},
						},
					},
				},
			},
		},
	}
	m := machine.New(&s.Machine)
	m.Boot()
}
