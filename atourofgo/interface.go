package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe1(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	describe(i)
	//i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	var j interface{}
	describe1(j)

	j = 42
	describe1(j)

	j = "hello"
	describe1(j)
}
