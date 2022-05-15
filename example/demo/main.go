package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xsuners/machine"
	"github.com/xsuners/machine/database"
)

func main() {

	// init databases
	database.Init(s.Databases...)
	defer database.Close()

	// init machine
	m := machine.New(&s.Machine)

	// d, _ := json.Marshal(s)
	// fmt.Println(string(d))

	// boot machine
	m.Boot()
}
