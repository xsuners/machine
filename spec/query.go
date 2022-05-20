package spec

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xsuners/machine/spec/types"
)

type Query struct {
	Type   types.QT
	Prop   string
	Kind   types.PK
	Values []any
}

func (s *Query) Set(path string, data any, op ...string) error {
	parts := strings.Split(path, ".")
	if len(parts) != 1 {
		return fmt.Errorf("path %s invalid", strings.Join(parts, "."))
	}
	var ok bool
	switch parts[0] {
	case "type":
		s.Type, ok = data.(types.QT)
		if !ok {
			return errors.New("set query: data not qt")
		}
	case "prop":
		s.Prop, ok = data.(string)
		if !ok {
			return errors.New("set query: data not string")
		}
	case "kind":
		s.Kind, ok = data.(types.PK)
		if !ok {
			return errors.New("set query: data not pk")
		}
	case "values":
		if len(op) > 0 {
			if op[0] == "append" {
				s.Values = append(s.Values, data)
				return nil
			} else {
				return errors.New("set query: path too long")
			}
		}
		s.Values, ok = data.([]any)
		if !ok {
			return errors.New("set query: data not []any")
		}
	default:
		return errors.New("set query: no feild " + parts[0])
	}
	return nil
}
