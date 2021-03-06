package in

import (
	"fmt"
	"strings"
)

type Object struct {
	Database string
	Table    string
	Id       int64
	Props    []*Prop
}

func (o *Object) key() string {
	return fmt.Sprintf("%s-%s-%d", o.Database, o.Table, o.Id)
}

func (o *Object) Get(path string) (any, bool) {
	ss := strings.SplitN(path, ".", 2)
	switch len(ss) {
	case 1:
		switch ss[0] {
		case "":
			return o, true
		case "database":
			return o.Database, true
		case "table":
			return o.Table, true
		case "id":
			return o.Id, true
		case "props":
			return o.Props, true
		}
	case 2:
		if ss[0] == "props" {
			return Props(o.Props).Get(ss[1])
		}
	}
	return nil, false
}

type Objects []*Object

func (ps Objects) Get(path string) (any, bool) {
	ss := strings.SplitN(path, ".", 2)
	switch len(ss) {
	case 1:
		if ss[0] == "" {
			return ps, true
		}
		for _, p := range ps {
			if ss[0] == p.key() {
				return p, true
			}
		}
	case 2:
		for _, p := range ps {
			if ss[0] == p.key() {
				return p.Get(ss[1])
			}
		}
	}
	return nil, false
}
