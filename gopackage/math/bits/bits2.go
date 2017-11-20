package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var a uint = 31
	fmt.Printf("bits.OnesCount(%d) = %d\n", a, bits.OnesCount(a))

	a++
	fmt.Printf("bits.OnesCount(%d) = %d\n", a, bits.OnesCount(a))
}
