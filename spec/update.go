package spec

import (
	"errors"
	"fmt"
	"strings"
)

type Update struct {
	Database string
	Table    string
	Queries  []*Query
	Props    []*Prop
}

func (u *Update) Get(path string) (any, bool) {
	ss := strings.SplitN(path, ".", 2)
	switch len(ss) {
	case 1:
		switch ss[0] {
		case "":
			return u, true
		case "database":
			return u.Database, true
		case "table":
			return u.Table, true
		case "queries":
			return u.Queries, true
		case "props":
			return u.Props, true
		}
	case 2:
		switch ss[0] {
		case "queries":
			return Queries(u.Queries).Get(ss[1])
		case "props":
			return Props(u.Props).Get(ss[1])
		}
	}
	return nil, false
}

func (u *Update) Set(path string, data any, op ...string) error {
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
		u.Database, ok = data.(string)
		if !ok {
			return errors.New("data type not string")
		}
	case "table":
		if len(parts) > 1 {
			return errors.New("path too long")
		}
		u.Table, ok = data.(string)
		if !ok {
			return errors.New("data type not string")
		}
	case "queries":
		if len(parts) == 1 {
			u.Queries, ok = data.([]*Query)
			if !ok {
				return errors.New("data type not int")
			}
			return nil
		}
		for _, q := range u.Queries {
			if q.Prop == parts[1] {
				return q.Set(parts[2], data, op...)
			}
		}
		return fmt.Errorf("set u: prop %s not found", path)
	case "props":
		if len(parts) == 1 {
			u.Props, ok = data.([]*Prop)
			if !ok {
				return errors.New("data type not int")
			}
			return nil
		}
		for _, q := range u.Props {
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
