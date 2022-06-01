package context

import (
	"fmt"
	"strings"
)

type memory map[string]any

func (m memory) Get(path string) (any, bool) {
	parts := strings.Split(path, ".")
	t := &m
	for i, part := range parts {
		if i == len(parts)-1 {
			data, ok := (*t)[part]
			return data, ok
		}
		d, ok := (*t)[part].(memory)
		if !ok {
			fmt.Println("zhe", path)
			return nil, false
		}
		t = &d
	}
	return nil, false
}

func (m memory) Set(path string, data any, op ...string) error {
	parts := strings.Split(path, ".")
	if m == nil {
		m = make(memory)
	}
	t := &m
	for i, part := range parts {
		if i == len(parts)-1 {
			(*t)[part] = data
			return nil
		}
		if _, ok := (*t)[part]; !ok {
			o := make(memory)
			(*t)[part] = o
			t = &o
			continue
		} else {
			d, ok := (*t)[part].(memory)
			if !ok {
				return fmt.Errorf("path %s has not map value", path)
			}
			t = &d
		}
	}
	return nil
}
