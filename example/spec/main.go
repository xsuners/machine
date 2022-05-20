package main

import (
	"encoding/json"
	"fmt"

	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/machine/spec/types"
)

func main() {
	// list()
	update()
	// get()
}

func get() {
	ctx := context.New()
	ctx.In = spec.In{
		Event: spec.Event{
			Props: []*spec.Prop{
				{
					Name:  "id",
					Value: 100000,
				},
			},
		},
	}
	fmt.Println(ctx.Get("in.event.props"))
	fmt.Println(ctx.Get("in.event.props.id"))
	fmt.Println(ctx.Get("in.event.props.id.value"))
}

func update() {
	ctx := context.New()
	ctx.In = spec.In{
		Update: spec.Update{
			Queries: []*spec.Query{
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

func list() {
	var in spec.In
	err := in.Set("list.database", "machine")
	if err != nil {
		panic(err)
	}
	err = in.Set("list.table", "user")
	if err != nil {
		panic(err)
	}
	err = in.Set("list.selects", []*spec.Select{
		{
			Prop: "name",
		},
	})
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
