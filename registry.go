package machine

type Builder func(map[string]interface{}) Node

var registry = make(map[string]Builder)

func Register(name string, builder Builder) {
	registry[name] = builder
}
