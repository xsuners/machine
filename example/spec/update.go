package main

import (
	"encoding/json"
	"fmt"

	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/machine/spec/in"
)

func getupdate() {
	d := in.Update{
		Database: "machine",
		Table:    "user",
		Queries: []*in.Query{
			{
				Type:   spec.Eq,
				Kind:   spec.Int,
				Prop:   "id",
				Values: []any{1},
			},
		},
		Props: []*in.Prop{
			{
				Name:  "id",
				Value: 1,
				Kind:  spec.Int,
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

	fmt.Println(d.Get("props"))
	fmt.Println(d.Get("props.id"))
	fmt.Println(d.Get("props.id.kind"))
	fmt.Println(d.Get("props.id.name"))
	fmt.Println(d.Get("props.id.value"))
}

func setupdate() {
	ctx := context.New()
	ctx.In = in.In{
		Update: in.Update{
			Queries: []*in.Query{
				{
					Prop: "id",
				},
			},
		},
	}
	err := ctx.Set("in.update.database", "machine")
	if err != nil {
		panic(err)
	}
	err = ctx.Set("in.update.table", "user")
	if err != nil {
		panic(err)
	}
	fmt.Println(ctx.Get("in.update.queries.id"))
	err = ctx.Set("in.update.queries.id.values", 20, "append")
	if err != nil {
		panic(err)
	}
	// err = ctx.Set("in.update.queries", []*spec.Query{
	// 	{
	// 		Prop: "name",
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// err = ctx.Set("in.update.queries.name.kind", types.String)
	// if err != nil {
	// 	panic(err)
	// }
	d, _ := json.Marshal(ctx.In)
	fmt.Println(string(d))
}
