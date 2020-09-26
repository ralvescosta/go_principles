package main

import (
	"fmt"
	"module/output_module"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Main")
	output_module.Output()

	error := checkmail.ValidateFormat("rafael.rac.mg@gmail.com")
	fmt.Println(error)
}
