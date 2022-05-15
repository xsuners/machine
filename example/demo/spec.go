package main

import (
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/machine/spec/data"
)

var s = &spec.Spec{
	Databases: []*data.Database{
		{
			Host:     "127.0.0.1",
			Port:     3306,
			Driver:   "mysql",
			Username: "root",
			Password: "123456",
			Name:     "group",
		},
	},
	Machine: spec.Machine{
		Rpcs: []*spec.Rpc{
			{
				Path: "test",
				Root: &spec.Node{
					Type:  "sequence",
					Name:  "s1",
					Props: make(map[string]any),
					Children: []*spec.Node{
						{
							Type:  "create",
							Name:  "c1",
							Props: make(map[string]any),
						},
						{
							Type:  "record",
							Name:  "r2",
							Props: make(map[string]any),
						},
					},
				},
			},
			{
				Path: "list",
				Root: &spec.Node{
					Type:  "sequence",
					Name:  "s1",
					Props: make(map[string]any),
					Children: []*spec.Node{
						{
							Type:  "list",
							Name:  "l1",
							Props: make(map[string]any),
						},
					},
				},
			},
			{
				Path: "update",
				Root: &spec.Node{
					Type:  "sequence",
					Name:  "s1",
					Props: make(map[string]any),
					Children: []*spec.Node{
						{
							Type:  "update",
							Name:  "u1",
							Props: make(map[string]any),
						},
					},
				},
			},
		},
	},
}
