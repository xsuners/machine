package spec

import "encoding/json"

func Load(raw []byte) *Spec {
	s := new(Spec)
	err := json.Unmarshal(raw, s)
	if err != nil {
		panic(err)
	}
	return s
}
