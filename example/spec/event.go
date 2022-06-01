package main

import (
	"fmt"

	"github.com/xsuners/machine/spec"
	"github.com/xsuners/machine/spec/types"
)

func getevent() {
	event := &spec.Event{
		Database: "dd",
		Table:    "tt",
		Id:       1,
		Props: []*spec.Prop{
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
	}
	fmt.Println(event.Get("database"))
	fmt.Println(event.Get("table"))
	fmt.Println(event.Get("id"))
	fmt.Println(event.Get("props.id"))
	fmt.Println(event.Get("props.id.name"))
	fmt.Println(event.Get("props.id.value"))
	fmt.Println(event.Get("props.id.kind"))
}
