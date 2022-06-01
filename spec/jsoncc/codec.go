package jsoncc

import "github.com/xsuners/machine/spec"

type jsoncc struct{}

func New() spec.Codec {
	return &jsoncc{}
}

func (j *jsoncc) Marshal(*spec.Spec) ([]byte, error)
func (j *jsoncc) Unmarshal([]byte, *spec.Spec) error
func (j *jsoncc) Name() string
