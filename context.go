package machine

import (
	"github.com/nats-io/nats.go"
)

type Context struct {
	machine *Machine
	memory  *Memory
	message *nats.Msg
	order   interface{}
}
