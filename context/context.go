package context

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/nats-io/nats.go"
)

type Context struct {
	order   interface{}
	request *http.Request
	message *nats.Msg
	mu      sync.RWMutex
	m       map[string]interface{}
}

func New(ins ...interface{}) *Context {
	ctx := new(Context)
	for _, in := range ins {
		switch m := in.(type) {
		case *nats.Msg:
			ctx.message = m
			var data interface{}
			err := json.Unmarshal(m.Data, &data)
			if err != nil {
				return ctx
			}
			err = ctx.Set("message", data)
			if err != nil {
				return ctx
			}
		case *http.Request:
			ctx.request = m
			if m.Body == nil {
				return ctx
			}
			decoder := json.NewDecoder(m.Body)
			defer m.Body.Close()
			var data interface{}
			err := decoder.Decode(&data)
			if err != nil {
				return ctx
			}
			err = ctx.Set("body", data)
			if err != nil {
				return ctx
			}
		}
	}
	return ctx
}

func (ctx *Context) Get(path string) (interface{}, bool) {
	parts := strings.Split(path, ".")
	ctx.mu.RLock()
	defer ctx.mu.RUnlock()
	t := &ctx.m
	if parts[0] == "header" || parts[0] == "query" {
		if len(parts[1]) < 0 {
			return nil, false
		}
		var v string
		switch parts[0] {
		case "header":
			v = ctx.request.Header.Get(parts[1])
		case "query":
			v = ctx.request.URL.Query().Get(parts[1])
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

func (ctx *Context) Set(path string, data interface{}) error {
	parts := strings.Split(path, ".")
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	if ctx.m == nil {
		ctx.m = make(map[string]interface{})
	}
	t := &ctx.m
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

func (ctx *Context) Bytes(path string) ([]byte, error) {
	m, ok := ctx.Get(path)
	if !ok {
		return nil, fmt.Errorf("path %s no value", path)
	}
	ctx.mu.RLock()
	defer ctx.mu.RUnlock()
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return data, nil
}
