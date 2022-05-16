package spec

import (
	"github.com/xsuners/machine/spec/types"
)

type Prop struct {
	Name  string
	Kind  types.PK
	Value any
}

type Object struct {
	Database string
	Table    string
	Props    []*Prop
}

type Create struct {
	Objects []*Object
}

type Query struct {
	Type   types.QT
	Prop   string
	Kind   types.PK
	Values []any
}

type Select struct {
	Prop string
	Kind types.PK
}

type List struct {
	Database string
	Table    string
	Selects  []*Select
	Queries  []*Query
	Page     int
	Size     int
}

type Update struct {
	Database string
	Table    string
	Queries  []*Query
	Props    []*Prop
}

type Delete struct {
	Database string
	Table    string
	Queries  []*Query
}

type Event struct {
	Database string
	Table    string
	Id       int64
	Props    []*Prop
}

type Message struct {
	List   List
	Create Create
	Update Update
	Delete Delete
	Event  Event
}
