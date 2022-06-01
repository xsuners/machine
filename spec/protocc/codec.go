package protocc

import "github.com/xsuners/machine/spec"

type protocc struct{}

func New() spec.Codec {
	return &protocc{}
}

func (j *protocc) Marshal(*spec.Spec) ([]byte, error)
func (j *protocc) Unmarshal([]byte, *spec.Spec) error
func (j *protocc) Name() string
