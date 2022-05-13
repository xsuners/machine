package machine

type Result string

const (
	Success Result = "success"
	Failure Result = "failure"
	Running Result = "running"
)

type Node interface {
	Id() string
	Exec(ctx *Context) (Result, error)
	Children(...Node) map[string]Node
}
