package main

import (
	"github.com/eugenewyj/go-sample/goinaction/ch05/ch0504/entities"
	"fmt"
)

func main() {
	u := entities.User7_1{
		Name: "Bill",
		Email: "bill@email.com",
	}

	fmt.Printf("User7_1: %v\n", u)
}
