package spec

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xsuners/machine/spec/types"
)

type Prop struct {
	Name  string
	Kind  types.PK
	Value any
}

func (prop *Prop) Set(data any, paths ...string) error {
	if len(paths) != 1 {
		return fmt.Errorf("path %s invalid", strings.Join(paths, "."))
	}
	var ok bool
	switch paths[0] {
	case "name":
		prop.Name, ok = data.(string)
		if !ok {
			return errors.New("set prop: data not string")
		}
	case "kind":
		prop.Kind, ok = data.(types.PK)
		if !ok {
			return errors.New("set prop: data not pk")
		}
	case "value":
		prop.Value = data
	default:
		return errors.New("set prop: no feild " + paths[0])
	}
	return nil
}
