package main

import (
	"encoding/json"
	"fmt"

	"github.com/xsuners/machine/spec"
	"github.com/xsuners/machine/spec/types"
)

func main() {
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
