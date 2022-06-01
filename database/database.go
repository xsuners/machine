package database

import (
	"github.com/xsuners/machine/spec"
	"github.com/xsuners/msql.v2"
)

var registry = make(map[string]msql.DB)
var cs []func()

func Init(databases ...*spec.Database) {
	for _, db := range databases {
		d, c, err := msql.New(
			msql.IP(db.Host),
			msql.Port(db.Port),
			msql.Username(db.Username),
			msql.Password(db.Password),
			msql.Driver(db.Driver),
			msql.Name(db.Name),
		)
		if err != nil {
			panic(err)
		}
		registry[db.Name] = d
		cs = append(cs, c)
	}
}

func Close() {
	for _, c := range cs {
		c()
	}
}

func Fetch(name string) msql.DB {
	db, ok := registry[name]
	if ok {
		return db
	} else {
		return &msql.Empty{}
	}
}
