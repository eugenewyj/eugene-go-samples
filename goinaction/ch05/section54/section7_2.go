package main

import (
	"github.com/eugenewyj/go-sample/goinaction/ch05/section54/entities"
	"fmt"
)

func main() {
	a := entities.Admin7_2{
		Rights: 10,
	}

	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("Admin_2: %v\n", a)
}
