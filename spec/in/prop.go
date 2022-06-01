package in

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xsuners/machine/spec"
)

type Prop struct {
	Name  string
	Kind  spec.PK
	Value any
}

func (p *Prop) Get(path string) (any, bool) {
	switch path {
	case "":
		return p, true
	case "kind":
		return p.Kind, true
	case "name":
		return p.Name, true
	case "value":
		return p.Value, true
	}
	return nil, false
}

func (p *Prop) Set(path string, data any) error {
	parts := strings.SplitN(path, ".", 2)
	if len(parts) != 1 {
		return fmt.Errorf("path %s invalid", strings.Join(parts, "."))
	}
	var ok bool
	switch parts[0] {
	case "name":
		p.Name, ok = data.(string)
		if !ok {
			return errors.New("set prop: data not string")
		}
	case "kind":
		p.Kind, ok = data.(spec.PK)
		if !ok {
			return errors.New("set prop: data not pk")
		}
	case "value":
		p.Value = data
	default:
		return errors.New("set prop: no feild " + parts[0])
	}
	return nil
}

type Props []*Prop

func (ps Props) Get(path string) (any, bool) {
	ss := strings.SplitN(path, ".", 2)
	switch len(ss) {
	case 1:
		if ss[0] == "" {
			return ps, true
		}
		for _, p := range ps {
			if ss[0] == p.Name {
				return p, true
			}
		}
	case 2:
		for _, p := range ps {
			if ss[0] == p.Name {
				return p.Get(ss[1])
			}
		}
	}
	return nil, false
}
