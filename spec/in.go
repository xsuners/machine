package spec

import (
	"errors"
	"fmt"
	"strings"
)

type In struct {
	List   List
	Create Create
	Update Update
	Delete Delete
	Event  Event
}

func (in *In) Set(path string, data any, op ...string) error {
	parts := strings.SplitN(path, ".", 2)
	if len(parts) < 1 {
		return errors.New("path invalid")
	}
	var ok bool
	switch parts[0] {
	case "list":
		if len(parts) > 1 {
			return in.List.Set(parts[1], data)
		}
		if in.List, ok = data.(List); !ok {
			return fmt.Errorf("data: %T not match %T", data, in.List)
		}
	// case "create":
	// 	if len(parts) == 1 {
	// 		in.Create, ok = data.(Create)
	// 		if !ok {
	// 			return fmt.Errorf("data: %T not match %T", data, in.Create)
	// 		}
	// 	} else {
	// 		return in.Create.Set(data, parts[1:]...)
	// 	}
	case "update":
		if len(parts) > 1 {
			return in.Update.Set(parts[1], data, op...)
		}
		if in.Update, ok = data.(Update); !ok {
			return fmt.Errorf("data: %T not match %T", data, in.Update)
		}
	case "delete":
		if len(parts) > 1 {
			return in.Delete.Set(parts[1], data)
		}
		if in.Delete, ok = data.(Delete); !ok {
			return fmt.Errorf("data: %T not match %T", data, in.Delete)
		}
	case "event":
		if len(parts) > 1 {
			return in.Event.Set(parts[1], data)
		}
		if in.Event, ok = data.(Event); !ok {
			return fmt.Errorf("data: %T not match %T", data, in.Event)
		}
	}
	return nil
}
