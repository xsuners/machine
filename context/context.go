package context

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/xsuners/machine/spec/in"
)

type Context struct {
	context.Context
	mu sync.RWMutex

	request *http.Request
	message *nats.Msg

	m  memory
	In in.In
}

func New(ins ...any) *Context {
	ctx := &Context{
		Context: context.Background(),
	}
	for _, in := range ins {
		switch m := in.(type) {
		case *nats.Msg:
			ctx.message = m
			var data any
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
			var data any
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

func (ctx *Context) Get(path string) (any, bool) {
	ctx.mu.RLock()
	defer ctx.mu.RUnlock()

	ss := strings.SplitN(path, ".", 2)
	switch len(ss) {
	case 1:
		switch ss[0] {
		case "header":
			return ctx.request.Header, true
		case "query":
			return ctx.request.URL.Query(), true
		case "in":
			return ctx.In, true
		}
	case 2:
		switch ss[0] {
		case "header":
			return ctx.request.Header.Get(ss[1]), true
		case "query":
			return ctx.request.URL.Query().Get(ss[1]), true
		case "in":
			return ctx.In.Get(ss[1])
		}
	}

	return ctx.m.Get(path)
}

func (ctx *Context) Set(path string, data any, op ...string) error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	ps := strings.SplitN(path, ".", 2)
	switch ps[0] {
	case "in":
		return ctx.In.Set(ps[1], data, op...)
	}

	return ctx.m.Set(path, data, op...)
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
