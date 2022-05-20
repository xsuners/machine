package context

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/xsuners/machine/spec"
)

type Context struct {
	context.Context

	request *http.Request
	message *nats.Msg
	mu      sync.RWMutex
	m       map[string]any

	In spec.In
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

	parts := strings.Split(path, ".")

	if len(parts) < 1 {
		return nil, false
	}
	switch parts[0] {
	case "header":
		if len(parts) == 2 {
			return ctx.request.Header.Get(parts[1]), true
		} else {
			return nil, false
		}
	case "query":
		if len(parts) == 2 {
			return ctx.request.URL.Query().Get(parts[1]), true
		} else {
			return nil, false
		}
	case "in":
		return ctx.getin(path)
	}

	t := &ctx.m
	for i, part := range parts {
		if i == len(parts)-1 {
			data, ok := (*t)[part]
			return data, ok
		}
		d, ok := (*t)[part].(map[string]any)
		if !ok {
			fmt.Println("zhe", path)
			return nil, false
		}
		t = &d
	}
	return nil, false
}

func (ctx *Context) getin(path string) (any, bool) {
	parts := strings.Split(path, ".")

	if len(parts) == 1 {
		return ctx.In, true
	}

	switch parts[1] {
	case "create":
		if len(parts) == 2 {
			return ctx.In.Create, true
		}
		// TODO
	case "update":
		if len(parts) == 2 {
			return ctx.In.Update, true
		}
		switch parts[2] {
		case "database":
			if len(parts) == 3 {
				return ctx.In.Update.Database, true
			} else {
				return nil, false
			}
		case "table":
			if len(parts) == 3 {
				return ctx.In.Update.Table, true
			} else {
				return nil, false
			}
		case "queries":
			if len(parts) == 3 {
				return ctx.In.Update.Queries, true
			}
			for _, query := range ctx.In.Update.Queries {
				if query.Prop == parts[3] {
					if len(parts) == 4 {
						return query, true
					}
					switch parts[4] {
					case "type":
						if len(parts) == 5 {
							return query.Type, true
						}
						return nil, false
					case "prop":
						if len(parts) == 5 {
							return query.Prop, true
						}
						return nil, false
					case "kind":
						if len(parts) == 5 {
							return query.Kind, true
						}
						return nil, false
					case "values":
						if len(parts) == 5 {
							return query.Values, true
						}
						return nil, false
					}
					return nil, false
				}
			}
			return nil, false
		case "props":
			if len(parts) == 3 {
				return ctx.In.Update.Props, true
			}
			for _, prop := range ctx.In.Update.Props {
				if prop.Name == parts[3] {
					if len(parts) == 4 {
						return prop, true
					}
					switch parts[4] {
					case "name":
						if len(parts) == 5 {
							return prop.Name, true
						} else {
							return nil, false
						}
					case "kind":
						if len(parts) == 5 {
							return prop.Kind, true
						} else {
							return nil, false
						}
					case "value":
						if len(parts) == 5 {
							return prop.Value, true
						} else {
							return nil, false
						}
					}
					return nil, false
				}
			}
			return nil, false
		}
	case "delete":
		if len(parts) == 2 {
			return ctx.In.Delete, true
		}
		switch parts[2] {
		case "database":
			if len(parts) == 3 {
				return ctx.In.Delete.Database, true
			} else {
				return nil, false
			}
		case "table":
			if len(parts) == 3 {
				return ctx.In.Delete.Table, true
			} else {
				return nil, false
			}
		case "queries":
			if len(parts) == 3 {
				return ctx.In.Delete.Queries, true
			}
			for _, query := range ctx.In.Delete.Queries {
				if query.Prop == parts[3] {
					if len(parts) == 4 {
						return query, true
					}
					switch parts[4] {
					case "type":
						if len(parts) == 5 {
							return query.Type, true
						}
						return nil, false
					case "prop":
						if len(parts) == 5 {
							return query.Prop, true
						}
						return nil, false
					case "kind":
						if len(parts) == 5 {
							return query.Kind, true
						}
						return nil, false
					case "values":
						if len(parts) == 5 {
							return query.Values, true
						}
						return nil, false
					}
					return nil, false
				}
			}
			return nil, false
		}
	case "list":
		if len(parts) == 2 {
			return ctx.In.List, true
		}
		switch parts[2] {
		case "database":
			if len(parts) == 3 {
				return ctx.In.List.Database, true
			} else {
				return nil, false
			}
		case "table":
			if len(parts) == 3 {
				return ctx.In.List.Table, true
			} else {
				return nil, false
			}
		case "queries":
			if len(parts) == 3 {
				return ctx.In.List.Queries, true
			}
			for _, query := range ctx.In.List.Queries {
				if query.Prop == parts[3] {
					if len(parts) == 4 {
						return query, true
					}
					switch parts[4] {
					case "type":
						if len(parts) == 5 {
							return query.Type, true
						}
						return nil, false
					case "prop":
						if len(parts) == 5 {
							return query.Prop, true
						}
						return nil, false
					case "kind":
						if len(parts) == 5 {
							return query.Kind, true
						}
						return nil, false
					case "values":
						if len(parts) == 5 {
							return query.Values, true
						}
						return nil, false
					}
					return nil, false
				}
			}
			return nil, false
		}
	case "event":
		if len(parts) == 2 {
			return ctx.In.Event, true
		}
		switch parts[2] {
		case "database":
			if len(parts) == 3 {
				return ctx.In.Event.Database, true
			} else {
				return nil, false
			}
		case "table":
			if len(parts) == 3 {
				return ctx.In.Event.Table, true
			} else {
				return nil, false
			}
		case "id":
			if len(parts) == 3 {
				return ctx.In.Event.Id, true
			} else {
				return nil, false
			}
		case "props":
			if len(parts) == 3 {
				return ctx.In.Event.Props, true
			}
			for _, prop := range ctx.In.Event.Props {
				if prop.Name == parts[3] {
					if len(parts) == 4 {
						return prop, true
					}
					switch parts[4] {
					case "name":
						if len(parts) == 5 {
							return prop.Name, true
						} else {
							return nil, false
						}
					case "kind":
						if len(parts) == 5 {
							return prop.Kind, true
						} else {
							return nil, false
						}
					case "value":
						if len(parts) == 5 {
							return prop.Value, true
						} else {
							return nil, false
						}
					}
					return nil, false
				}
			}
			return nil, false
		}
	}
	return nil, false
}

func (ctx *Context) Set(path string, data any, op ...string) error {
	parts := strings.Split(path, ".")
	nparts := strings.SplitN(path, ".", 2)
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	switch parts[0] {
	case "in":
		return ctx.In.Set(nparts[1], data, op...)
	}

	if ctx.m == nil {
		ctx.m = make(map[string]any)
	}
	t := &ctx.m
	for i, part := range parts {
		if i == len(parts)-1 {
			(*t)[part] = data
			return nil
		}
		if _, ok := (*t)[part]; !ok {
			o := make(map[string]any)
			(*t)[part] = o
			t = &o
			continue
		} else {
			d, ok := (*t)[part].(map[string]any)
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
