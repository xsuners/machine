package spec

import (
	"errors"
)

type List struct {
	Database string
	Table    string
	Selects  []*Select
	Queries  []*Query
	Page     int
	Size     int
}

func (list *List) Set(data any, paths ...string) error {
	if len(paths) < 1 {
		return errors.New("invalid paths")
	}
	var ok bool
	switch paths[0] {
	case "database":
		if len(paths) > 1 {
			return errors.New("path too long")
		}
		list.Database, ok = data.(string)
		if !ok {
			return errors.New("data type not int")
		}
	case "table":
		if len(paths) > 1 {
			return errors.New("path too long")
		}
		list.Table, ok = data.(string)
		if !ok {
			return errors.New("data type not int")
		}
	case "selects":
		if len(paths) == 1 {
			list.Selects, ok = data.([]*Select)
			if !ok {
				return errors.New("data type not int")
			}
			return nil
		}
		for _, s := range list.Selects {
			if s.Prop == paths[1] {
				return s.Set(data, paths[2:]...)
			}
		}
	case "queries":
		if len(paths) == 1 {
			list.Queries, ok = data.([]*Query)
			if !ok {
				return errors.New("data type not int")
			}
		}
		for _, q := range list.Queries {
			if q.Prop == paths[1] {
				return q.Set(data, paths[2:]...)
			}
		}
	case "page":
		if len(paths) > 1 {
			return errors.New("path too long")
		}
		list.Page, ok = data.(int)
		if !ok {
			return errors.New("data type not int")
		}
	case "size":
		if len(paths) > 1 {
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
