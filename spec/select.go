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

func (s *Select) Set(data any, paths ...string) error {
	if len(paths) != 1 {
		return fmt.Errorf("path %s invalid", strings.Join(paths, "."))
	}
	var ok bool
	switch paths[0] {
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
		return errors.New("set select: no feild " + paths[0])
	}
	return nil
}
