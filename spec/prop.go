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

func (prop *Prop) Set(path string, data any) error {
	parts := strings.SplitN(path, ".", 2)
	if len(parts) != 1 {
		return fmt.Errorf("path %s invalid", strings.Join(parts, "."))
	}
	var ok bool
	switch parts[0] {
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
		return errors.New("set prop: no feild " + parts[0])
	}
	return nil
}
