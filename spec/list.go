package spec

import (
	"errors"
	"fmt"
	"strings"
)

type List struct {
	Database string
	Table    string
	Selects  []*Select
	Queries  []*Query
	Page     int
	Size     int
}

func (list *List) Set(path string, data any) error {
	parts := strings.SplitN(path, ".", 3)
	if len(parts) < 1 {
		return errors.New("invalid paths")
	}
	var ok bool
	switch parts[0] {
	case "database":
		if len(parts) > 1 {
			return errors.New("path too long")
		}
		list.Database, ok = data.(string)
		if !ok {
			return errors.New("data type not string")
		}
	case "table":
		if len(parts) > 1 {
			return errors.New("path too long")
		}
		list.Table, ok = data.(string)
		if !ok {
			return errors.New("data type not string")
		}
	case "selects":
		if len(parts) == 1 {
			list.Selects, ok = data.([]*Select)
			if !ok {
				return errors.New("data type not []*Select")
			}
			return nil
		}
		for _, s := range list.Selects {
			if s.Prop == parts[1] {
				return s.Set(parts[2], data)
			}
		}
		return fmt.Errorf("set list: prop %s not found", path)
	case "queries":
		if len(parts) == 1 {
			list.Queries, ok = data.([]*Query)
			if !ok {
				return errors.New("data type not int")
			}
			return nil
		}
		for _, q := range list.Queries {
			if q.Prop == parts[1] {
				return q.Set(parts[2], data)
			}
		}
		return fmt.Errorf("set list: prop %s not found", path)
	case "page":
		if len(parts) > 1 {
			return errors.New("path too long")
		}
		list.Page, ok = data.(int)
		if !ok {
			return errors.New("data type not int")
		}
	case "size":
		if len(parts) > 1 {
			return errors.New("path too long")
		}
		list.Size, ok = data.(int)
		if !ok {
			return errors.New("data type not int")
		}
	default:
		return errors.New("in no field in paths")
	}
	return nil
}
