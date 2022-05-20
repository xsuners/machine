package spec

import (
	"errors"
	"fmt"
	"strings"
)

type Create struct {
	Objects []*Object
}

func (u *Create) Set(path string, data any) error {
	parts := strings.SplitN(path, ".", 3)
	if len(parts) < 1 {
		return errors.New("invalid paths")
	}
	var ok bool
	switch parts[0] {
	case "objects":
		if len(parts) == 1 {
			u.Objects, ok = data.([]*Object)
			if !ok {
				return errors.New("data type not int")
			}
			return nil
		}
		// for _, q := range u.Objects {
		// 	if q.Prop == parts[1] {
		// 		return q.Set(parts[2], data)
		// 	}
		// }
		return fmt.Errorf("set u: prop %s not found", path)
	default:
		return errors.New("in no field in paths")
	}
	// return nil
}
