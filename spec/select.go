package spec

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xsuners/machine/spec/types"
)

type Select struct {
	Prop string
	Kind types.PK
}

func (c *Select) Get(path string) (any, bool) {
	ss := strings.SplitN(path, ".", 2)
	switch len(ss) {
	case 1:
		switch ss[0] {
		case "":
			return c, true
		case "prop":
			return c.Prop, true
		case "kind":
			return c.Kind, true
		}
	}
	return nil, false
}

func (s *Select) Set(path string, data any) error {
	parts := strings.Split(path, ".")
	if len(parts) != 1 {
		return fmt.Errorf("path %s invalid", strings.Join(parts, "."))
	}
	var ok bool
	switch parts[0] {
	case "prop":
		s.Prop, ok = data.(string)
		if !ok {
			return errors.New("set select: data not string")
		}
	case "kind":
		s.Kind, ok = data.(types.PK)
		if !ok {
			return errors.New("set select: data not pk")
		}
	default:
		return errors.New("set select: no feild " + parts[0])
	}
	return nil
}

type Selects []*Select

func (ps Selects) Get(path string) (any, bool) {
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
