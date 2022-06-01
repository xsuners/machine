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

func (d *Delete) Get(path string) (any, bool) {
	ss := strings.SplitN(path, ".", 2)
	switch len(ss) {
	case 1:
		switch ss[0] {
		case "":
			return d, true
		case "database":
			return d.Database, true
		case "table":
			return d.Table, true
		case "queries":
			return d.Queries, true
		}
	case 2:
		if ss[0] == "queries" {
			return Queries(d.Queries).Get(ss[1])
		}
	}
	return nil, false
}

func (d *Delete) Set(path string, data any) error {
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
		d.Database, ok = data.(string)
		if !ok {
			return errors.New("data type not string")
		}
	case "table":
		if len(parts) > 1 {
			return errors.New("path too long")
		}
		d.Table, ok = data.(string)
		if !ok {
			return errors.New("data type not string")
		}
	case "queries":
		if len(parts) == 1 {
			d.Queries, ok = data.([]*Query)
			if !ok {
				return errors.New("data type not int")
			}
			return nil
		}
		for _, q := range d.Queries {
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
