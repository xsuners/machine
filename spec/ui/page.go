package ui

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
