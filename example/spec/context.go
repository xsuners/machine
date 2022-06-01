package main

import (
	"fmt"

	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/spec"
)

func getctx() {
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
