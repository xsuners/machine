package spec

type Query struct {
	Types  []QT
	Prop   string
	Kind   PK
	Values []any
}

type Prop struct {
	Name   string
	Kind   PK
	Values []any
}

type PT string

const (
	List   PT = "list"
	Detail PT = "detail"
)

type Page struct {
	Type     PT
	Title    string
	Database string
	Table    string
	Select   []string
	Queries  []*Query
	Props    []*Prop
	Size     int
}
