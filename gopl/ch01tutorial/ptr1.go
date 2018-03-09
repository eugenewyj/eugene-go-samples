package main

import "fmt"

func main() {
	b := 1
	fmt.Printf("b = %d\n", b)
	var a = &b
	*a = 2
	fmt.Printf("b = %d, *a = %d, a = %d\n", b, *a, a)
}
