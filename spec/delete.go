package spec

type Delete struct {
	Database string
	Table    string
	Queries  []*Query
}

func (d *Delete) Set(data any, paths ...string) error {
	return nil
}
