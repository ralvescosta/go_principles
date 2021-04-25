package main

import (
	"fmt"
	database "study/src/database"
	module "study/src/module"
)

func main() {
	db := database.Database()

	m := module.Module(&db)
	some := m.DoSomething(1)

	fmt.Println(some)
}
