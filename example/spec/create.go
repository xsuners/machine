package main

import (
	"fmt"

	"github.com/xsuners/machine/spec/in"
	"github.com/xsuners/machine/spec/types"
)

func getcreate() {
	c := &in.Create{
		Objects: []*in.Object{
			{
				Database: "machine",
				Table:    "user",
				Id:       100000,
				Props: []*in.Prop{
					{
						Name:  "name",
						Value: "hello",
						Kind:  types.String,
					},
					{
						Name:  "id",
						Value: 1,
						Kind:  types.Int,
					},
				},
			},
		},
	}
	fmt.Println(c.Get(""))
	fmt.Println(c.Get("objects"))
	fmt.Println(c.Get("objects.machine-user-100000"))
	fmt.Println(c.Get("objects.machine-user-100000.database"))
	fmt.Println(c.Get("objects.machine-user-100000.table"))
	fmt.Println(c.Get("objects.machine-user-100000.id"))
	fmt.Println(c.Get("objects.machine-user-100000.props"))
	fmt.Println(c.Get("objects.machine-user-100000.props.name"))
	fmt.Println(c.Get("objects.machine-user-100000.props.name.name"))
	fmt.Println(c.Get("objects.machine-user-100000.props.name.value"))
	fmt.Println(c.Get("objects.machine-user-100000.props.name.kind"))
}
