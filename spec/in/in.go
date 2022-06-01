package in

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

func (c *In) Get(path string) (any, bool) {
	ss := strings.SplitN(path, ".", 2)
	switch len(ss) {
	case 1:
		switch ss[0] {
		case "":
			return c, true
		case "list":
			return c.List, true
		case "create":
			return c.Create, true
		case "update":
			return c.Update, true
		case "event":
			return c.Event, true
		}
	case 2:
		switch ss[0] {
		case "list":
			return c.List.Get(ss[1])
		case "create":
			return c.Create.Get(ss[1])
		case "update":
			return c.Update.Get(ss[1])
		case "event":
			return c.Event.Get(ss[1])
		}
	}
	return nil, false
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
	case "create":
		if len(parts) > 1 {
			return in.Create.Set(parts[1], data)
		}
		if in.Create, ok = data.(Create); !ok {
			return fmt.Errorf("data: %T not match %T", data, in.Create)
		}
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
