package data

type Database struct {
	Host     string
	Port     int
	Driver   string
	Username string
	Password string
	Name     string
	Tables   []*Table
}

type Table struct {
	Name    string
	Columns []*Column
}

type Column struct {
	Name    string
	Type    string
	Comment string
}
