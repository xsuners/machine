package ui

import (
	"github.com/xsuners/machine/spec/types"
)

type Query struct {
	Types  []types.QT
	Prop   string
	Kind   types.PK
	Values []any
}

type Prop struct {
	Name   string
	Kind   types.PK
	Values []any
}
