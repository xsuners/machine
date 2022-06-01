package main

import (
	"fmt"

	"github.com/xsuners/machine/spec/in"
	"github.com/xsuners/machine/spec/types"
)

func getdelete() {
	d := in.Delete{
		Database: "machine",
		Table:    "user",
		Queries: []*in.Query{
			{
				Type:   types.Eq,
				Kind:   types.Int,
				Prop:   "id",
				Values: []any{1},
			},
		},
	}
	fmt.Println(d.Get("database"))
	fmt.Println(d.Get("table"))
	fmt.Println(d.Get("queries"))
	fmt.Println(d.Get("queries.id"))
	fmt.Println(d.Get("queries.id.type"))
	fmt.Println(d.Get("queries.id.kind"))
	fmt.Println(d.Get("queries.id.prop"))
	fmt.Println(d.Get("queries.id.values"))
}
