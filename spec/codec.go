package spec

type Marshaler interface {
	Marshal(*Spec) ([]byte, error)
}

type Unmarshaler interface {
	Unmarshal([]byte, *Spec) error
}

type Codec interface {
	Marshaler
	Unmarshaler
	Name() string
}
