package spec

type Create struct {
	Objects []*Object
}

func (c *Create) Set(data any, paths ...string) error {
	return nil
}
