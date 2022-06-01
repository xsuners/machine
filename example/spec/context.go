package main

import (
	"fmt"

	"github.com/xsuners/machine/context"
	"github.com/xsuners/machine/spec/in"
)

func getctx() {
	ctx := context.New()
	ctx.In = in.In{
		Event: in.Event{
			Props: []*in.Prop{
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
