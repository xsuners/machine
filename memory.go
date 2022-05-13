package machine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type Memory struct {
	request *http.Request
	mu      sync.RWMutex
	m       map[string]interface{}
}

func newbb(request *http.Request) *Memory {
	m := &Memory{
		request: request,
	}
	if request.Body == nil {
		return m
	}
	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()
	var data interface{}
	err := decoder.Decode(&data)
	if err != nil {
		return m
	}
	err = m.Set("body", data)
	if err != nil {
		return m
	}
	return m
}

func (mem *Memory) Get(path string) (interface{}, bool) {
	parts := strings.Split(path, ".")
	mem.mu.RLock()
	defer mem.mu.RUnlock()
	t := &mem.m
	if parts[0] == "header" || parts[0] == "query" {
		if len(parts[1]) < 0 {
			return nil, false
		}
		var v string
		switch parts[0] {
		case "header":
			v = mem.request.Header.Get(parts[1])
		case "query":
			v = mem.request.URL.Query().Get(parts[1])
		}
		if len(v) < 1 {
			return nil, false
		} else {
			return v, true
		}
	}
	for i, part := range parts {
		if i == len(parts)-1 {
			data, ok := (*t)[part]
			return data, ok
		}
		d, ok := (*t)[part].(map[string]interface{})
		if !ok {
			fmt.Println("zhe", path)
			return nil, false
		}
		t = &d
	}
	return nil, false
}

func (mem *Memory) Set(path string, data interface{}) error {
	parts := strings.Split(path, ".")
	mem.mu.Lock()
	defer mem.mu.Unlock()
	if mem.m == nil {
		mem.m = make(map[string]interface{})
	}
	t := &mem.m
	for i, part := range parts {
		if i == len(parts)-1 {
			(*t)[part] = data
			return nil
		}
		if _, ok := (*t)[part]; !ok {
			o := make(map[string]interface{})
			(*t)[part] = o
			t = &o
			continue
		} else {
			d, ok := (*t)[part].(map[string]interface{})
			if !ok {
				return fmt.Errorf("path %s has not map value", path)
			}
			t = &d
		}
	}
	return nil
}

func (mem *Memory) Bytes(path string) ([]byte, error) {
	m, ok := mem.Get(path)
	if !ok {
		return nil, fmt.Errorf("path %s no value", path)
	}
	mem.mu.RLock()
	defer mem.mu.RUnlock()
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return data, nil
}
