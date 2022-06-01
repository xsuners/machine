package util

import (
	"fmt"
	"strings"

	"github.com/xsuners/machine/spec"
	"github.com/xsuners/machine/spec/in"
)

func Scan(ss ...*in.Select) (data map[string]any, vals []any) {
	data = make(map[string]any)
	var holders []any
	for _, s := range ss {
		var h any
		switch s.Kind {
		case spec.Int:
			h = int(0)
		case spec.Bool:
			h = false
		case spec.Float:
			h = float64(0.0)
		case spec.String:
			h = ""
		case spec.Enum:
			h = ""
		default:
			h = ""
		}
		data[s.Prop] = &h
		holders = append(holders, &h)
	}
	return data, holders
}

func Where(l []*in.Query, vals ...any) (string, []any) {
	// var vals []any
	var cons []string
	for _, query := range l {
		if len(query.Values) < 1 {
			continue
		}
		switch query.Type {
		case spec.Eq:
			cons = append(cons, query.Prop+" = ?")
			vals = append(vals, query.Values[0])
		case spec.Gt:
			cons = append(cons, query.Prop+" > ?")
			vals = append(vals, query.Values[0])
		case spec.Ge:
			cons = append(cons, query.Prop+" >= ?")
			vals = append(vals, query.Values[0])
		case spec.Lt:
			cons = append(cons, query.Prop+" < ?")
			vals = append(vals, query.Values[0])
		case spec.Le:
			cons = append(cons, query.Prop+" <= ?")
			vals = append(vals, query.Values[0])
		case spec.In:
			cons = append(cons, query.Prop+" in (?)")
			var vs []string
			for _, iv := range query.Values {
				vs = append(vs, fmt.Sprintf("%v", iv))
			}
			vals = append(vals, strings.Join(vs, ","))
		case spec.Ni:
			cons = append(cons, query.Prop+" not in (?)")
			var vs []string
			for _, iv := range query.Values {
				vs = append(vs, fmt.Sprintf("%v", iv))
			}
			vals = append(vals, strings.Join(vs, ","))
		}
	}
	return strings.Join(cons, " and "), vals
}

func Set(props []*in.Prop, vals ...any) (string, []any) {
	var keys []string
	for _, prop := range props {
		keys = append(keys, prop.Name+" = ?")
		vals = append(vals, prop.Value)
	}
	return strings.Join(keys, ","), vals
}
