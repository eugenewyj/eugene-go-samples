package main

import (
	"fmt"
	"math"
)

type XY struct {
	X, Y float64
}

func (v XY) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v XY) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *XY) ScalaPointer(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := XY{3, 4}
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	v.ScalaPointer(10)
	fmt.Println(v.Abs())

	var p = &v

	p.Scale(10)
	fmt.Println(p.Abs())

	p.ScalaPointer(10)
	fmt.Println(p.Abs())
}
