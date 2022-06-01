package main

import (
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/machine/spec/in"
)

var s = &spec.Spec{
	Databases: []*spec.Database{
		{
			Host:     "127.0.0.1",
			Port:     3306,
			Driver:   "mysql",
			Username: "root",
			Password: "123456",
			Name:     "machine",
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
		Mq: []*spec.Mq{
			{
				Subject: "goods.created",
				Root: &spec.Node{
					Type:  "sequence",
					Name:  "s1",
					Props: make(map[string]any),
					Children: []*spec.Node{
						{
							Type: "convert",
							Name: "c1",
							Props: map[string]any{
								"update": &in.Update{
									Database: "machine",
									Table:    "user",
									Queries: []*in.Query{
										{
											Type:   spec.Eq,
											Prop:   "id",
											Kind:   spec.Int,
											Values: []any{},
										},
									},
									Props: []*in.Prop{
										{
											Name:  "username",
											Kind:  spec.String,
											Value: "属性",
										},
									},
								},
								"statements": []string{
									"append in.update.queries.id.values in.event.props.id.value",
								},
							},
						},
						{
							Type: "logging",
							Name: "l1",
							Props: map[string]any{
								"database": "in.update.database",
								"queries":  "in.update.queries",
								"id":       "in.update.queries.id",
							},
						},
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
