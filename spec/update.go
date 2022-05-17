package spec

type Update struct {
	Database string
	Table    string
	Queries  []*Query
	Props    []*Prop
}

func (u *Update) Set(data any, paths ...string) error {
	return nil
}
