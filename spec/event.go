package spec

type Event struct {
	Database string
	Table    string
	Id       int64
	Props    []*Prop
}

func (e *Event) Set(data any, paths ...string) error {
	return nil
}
