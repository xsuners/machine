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

func (in *In) Set(path string, data any) error {
	parts := strings.Split(path, ".")
	if len(parts) < 1 {
		return errors.New("path invalid")
	}
	var ok bool
	switch parts[0] {
	case "list":
		if len(parts) == 1 {
			in.List, ok = data.(List)
			if !ok {
				return fmt.Errorf("data: %T not match %T", data, in.List)
			}
		} else {
			return in.List.Set(data, parts[1:]...)
		}
	case "create":
		if len(parts) == 1 {
			in.Create, ok = data.(Create)
			if !ok {
				return fmt.Errorf("data: %T not match %T", data, in.Create)
			}
		} else {
			return in.Create.Set(data, parts[1:]...)
		}
	case "update":
		if len(parts) == 1 {
			in.Update, ok = data.(Update)
			if !ok {
				return fmt.Errorf("data: %T not match %T", data, in.Update)
			}
		} else {
			return in.Update.Set(data, parts[1:]...)
		}
	case "delete":
		if len(parts) == 1 {
			in.Delete, ok = data.(Delete)
			if !ok {
				return fmt.Errorf("data: %T not match %T", data, in.Delete)
			}
		} else {
			return in.Delete.Set(data, parts[1:]...)
		}
	case "event":
		if len(parts) == 1 {
			in.Event, ok = data.(Event)
			if !ok {
				return fmt.Errorf("data: %T not match %T", data, in.Event)
			}
		} else {
			return in.Event.Set(data, parts[1:]...)
		}
	}
	return nil
}
