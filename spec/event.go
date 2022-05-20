package spec

import (
	"errors"
	"fmt"
	"strings"
)

type Event struct {
	Database string
	Table    string
	Id       int64
	Props    []*Prop
}

func (e *Event) Set(path string, data any) error {
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
		e.Database, ok = data.(string)
		if !ok {
			return errors.New("data type not string")
		}
	case "table":
		if len(parts) > 1 {
			return errors.New("path too long")
		}
		e.Table, ok = data.(string)
		if !ok {
			return errors.New("data type not string")
		}
	case "id":
		if len(parts) > 1 {
			return errors.New("path too long")
		}
		e.Id, ok = data.(int64)
		if !ok {
			return errors.New("data type not string")
		}
	case "props":
		if len(parts) == 1 {
			e.Props, ok = data.([]*Prop)
			if !ok {
				return errors.New("data type not int")
			}
			return nil
		}
		for _, q := range e.Props {
			if q.Name == parts[1] {
				return q.Set(parts[2], data)
			}
		}
		return fmt.Errorf("set u: prop %s not found", path)
	default:
		return errors.New("in no field in paths")
	}
	return nil
}
