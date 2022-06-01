package in

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xsuners/machine/spec"
)

type Query struct {
	Type   spec.QT
	Prop   string
	Kind   spec.PK
	Values []any
}

func (q *Query) Get(path string) (any, bool) {
	switch path {
	case "":
		return q, true
	case "type":
		return q.Type, true
	case "kind":
		return q.Kind, true
	case "prop":
		return q.Prop, true
	case "values":
		return q.Values, true
	}
	return nil, false
}

func (q *Query) Set(path string, data any, op ...string) error {
	parts := strings.Split(path, ".")
	if len(parts) != 1 {
		return fmt.Errorf("path %s invalid", strings.Join(parts, "."))
	}
	var ok bool
	switch parts[0] {
	case "type":
		q.Type, ok = data.(spec.QT)
		if !ok {
			return errors.New("set query: data not qt")
		}
	case "prop":
		q.Prop, ok = data.(string)
		if !ok {
			return errors.New("set query: data not string")
		}
	case "kind":
		q.Kind, ok = data.(spec.PK)
		if !ok {
			return errors.New("set query: data not pk")
		}
	case "values":
		if len(op) > 0 {
			if op[0] == "append" {
				q.Values = append(q.Values, data)
				return nil
			} else {
				return errors.New("set query: path too long")
			}
		}
		q.Values, ok = data.([]any)
		if !ok {
			return errors.New("set query: data not []any")
		}
	default:
		return errors.New("set query: no feild " + parts[0])
	}
	return nil
}

type Queries []*Query

func (ps Queries) Get(path string) (any, bool) {
	ss := strings.SplitN(path, ".", 2)
	switch len(ss) {
	case 1:
		if ss[0] == "" {
			return ps, true
		}
		for _, p := range ps {
			if ss[0] == p.Prop {
				return p, true
			}
		}
	case 2:
		for _, p := range ps {
			if ss[0] == p.Prop {
				return p.Get(ss[1])
			}
		}
	}
	return nil, false
}
