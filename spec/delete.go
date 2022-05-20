package spec

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
