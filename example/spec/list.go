package main

import (
	"encoding/json"
	"fmt"
	"go/types"

	"github.com/xsuners/machine/spec"
	"github.com/xsuners/machine/spec/in"
)

func getlist() {
	d := in.List{
		Database: "machine",
		Table:    "user",
		Page:     1,
		Size:     12,
		Selects: []*in.Select{
			{
				Prop: "id",
				Kind: spec.Int,
			},
		},
		Queries: []*in.Query{
			{
				Type:   spec.Eq,
				Kind:   spec.Int,
				Prop:   "id",
				Values: []any{1},
			},
		},
	}
	fmt.Println(d.Get("database"))
	fmt.Println(d.Get("table"))
	fmt.Println(d.Get("page"))
	fmt.Println(d.Get("size"))
	fmt.Println(d.Get("selects.id"))
	fmt.Println(d.Get("selects.id.prop"))
	fmt.Println(d.Get("selects.id.kind"))
	fmt.Println(d.Get("queries"))
	fmt.Println(d.Get("queries.id"))
	fmt.Println(d.Get("queries.id.type"))
	fmt.Println(d.Get("queries.id.kind"))
	fmt.Println(d.Get("queries.id.prop"))
	fmt.Println(d.Get("queries.id.values"))
}

func setlist() {
	var in in.In
	err := in.Set("list.database", "machine")
	if err != nil {
		panic(err)
	}
	err = in.Set("list.table", "user")
	if err != nil {
		panic(err)
	}
	err = in.Set("list.selects.name.kind", types.String)
	if err != nil {
		panic(err)
	}
	d, _ := json.Marshal(in)
	fmt.Println(string(d))
}
