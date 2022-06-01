package spec

type PK string

const (
	Int    PK = "int"
	Bool   PK = "bool"
	Float  PK = "float"
	String PK = "string"
	Enum   PK = "enum"
)

type QT string

const (
	Eq QT = "eq"
	Gt QT = "gt"
	Ge QT = "ge"
	Lt QT = "lt"
	Le QT = "le"
	In QT = "in"
	Ni QT = "ni"
)
