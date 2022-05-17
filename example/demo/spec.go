package main

import (
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/machine/spec/data"
	"github.com/xsuners/machine/spec/types"
)

var s = &spec.Spec{
	Databases: []*data.Database{
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
								"update": &spec.Update{
									Database: "machine",
									Table:    "user",
									Queries: []*spec.Query{
										{
											Type:   types.Eq,
											Prop:   "id",
											Kind:   types.Int,
											Values: []any{100000},
										},
									},
									Props: []*spec.Prop{
										{
											Name:  "username",
											Kind:  types.String,
											Value: "liule",
										},
									},
								},
								"statements": []any{
									&spec.Asign{
										From: "in.event.props.id",
										To:   "in.update.queries.id",
									},
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
