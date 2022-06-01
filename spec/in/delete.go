package in

import (
	"errors"
	"fmt"
	"strings"
)

type Delete struct {
	Database string
	Table    string
	Queries  []*Query
}

func (c *Delete) Get(path string) (any, bool) {
	ss := strings.SplitN(path, ".", 2)
	switch len(ss) {
	case 1:
		switch ss[0] {
		case "":
			return c, true
		case "database":
			return c.Database, true
		case "table":
			return c.Table, true
		case "queries":
			return c.Queries, true
		}
	case 2:
		if ss[0] == "queries" {
			return Queries(c.Queries).Get(ss[1])
		}
	}
	return nil, false
}

func (u *Delete) Set(path string, data any) error {
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
				return q.Set(parts[2], data)
			}
		}
		return fmt.Errorf("set u: prop %s not found", path)
	default:
		return errors.New("in no field in paths")
	}
	return nil
}
