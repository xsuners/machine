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
