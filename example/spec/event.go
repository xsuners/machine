package main

import (
	"fmt"

	"github.com/xsuners/machine/spec"
	"github.com/xsuners/machine/spec/in"
)

func getevent() {
	event := &in.Event{
		Database: "dd",
		Table:    "tt",
		Id:       1,
		Props: []*in.Prop{
			{
				Name:  "name",
				Value: "hello",
				Kind:  spec.String,
			},
			{
				Name:  "id",
				Value: 1,
				Kind:  spec.Int,
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
