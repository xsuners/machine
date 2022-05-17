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

func (s *Query) Set(data any, paths ...string) error {
	if len(paths) != 1 {
		return fmt.Errorf("path %s invalid", strings.Join(paths, "."))
	}
	var ok bool
	switch paths[0] {
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
		s.Values, ok = data.([]any)
		if !ok {
			return errors.New("set query: data not []any")
		}
	default:
		return errors.New("set query: no feild " + paths[0])
	}
	return nil
}
