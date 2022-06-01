package spec

import (
	"errors"
	"fmt"
	"strings"
)

type Create struct {
	Objects []*Object
}

func (c *Create) Get(path string) (any, bool) {
	ss := strings.SplitN(path, ".", 2)
	switch len(ss) {
	case 1:
		switch ss[0] {
		case "":
			return c, true
		case "objects":
			return c.Objects, true
		}
	case 2:
		if ss[0] == "objects" {
			return Objects(c.Objects).Get(ss[1])
		}
	}
	return nil, false
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
