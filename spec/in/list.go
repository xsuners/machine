package in

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

func (l *List) Get(path string) (any, bool) {
	ss := strings.SplitN(path, ".", 2)
	switch len(ss) {
	case 1:
		switch ss[0] {
		case "":
			return l, true
		case "database":
			return l.Database, true
		case "table":
			return l.Table, true
		case "page":
			return l.Page, true
		case "size":
			return l.Size, true
		case "queries":
			return l.Queries, true
		case "selects":
			return l.Selects, true
		}
	case 2:
		switch ss[0] {
		case "queries":
			return Queries(l.Queries).Get(ss[1])
		case "selects":
			return Selects(l.Selects).Get(ss[1])
		}
	}
	return nil, false
}

func (l *List) Set(path string, data any) error {
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
		l.Database, ok = data.(string)
		if !ok {
			return errors.New("data type not string")
		}
	case "table":
		if len(parts) > 1 {
			return errors.New("path too long")
		}
		l.Table, ok = data.(string)
		if !ok {
			return errors.New("data type not string")
		}
	case "selects":
		if len(parts) == 1 {
			l.Selects, ok = data.([]*Select)
			if !ok {
				return errors.New("data type not []*Select")
			}
			return nil
		}
		for _, s := range l.Selects {
			if s.Prop == parts[1] {
				return s.Set(parts[2], data)
			}
		}
		return fmt.Errorf("set list: prop %s not found", path)
	case "queries":
		if len(parts) == 1 {
			l.Queries, ok = data.([]*Query)
			if !ok {
				return errors.New("data type not int")
			}
			return nil
		}
		for _, q := range l.Queries {
			if q.Prop == parts[1] {
				return q.Set(parts[2], data)
			}
		}
		return fmt.Errorf("set list: prop %s not found", path)
	case "page":
		if len(parts) > 1 {
			return errors.New("path too long")
		}
		l.Page, ok = data.(int)
		if !ok {
			return errors.New("data type not int")
		}
	case "size":
		if len(parts) > 1 {
			return errors.New("path too long")
		}
		l.Size, ok = data.(int)
		if !ok {
			return errors.New("data type not int")
		}
	default:
		return errors.New("in no field in paths")
	}
	return nil
}
