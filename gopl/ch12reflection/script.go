package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"time"
	"github.com/eugenewyj/go-sample/gopl/ch12reflection/format"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	fmt.Printf("%T\n", 3)

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	t = v.Type()
	fmt.Println(t.String())

	x := v.Interface()
	i := x.(int)
	fmt.Printf("%d\n", i)

	var x1 int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(format.Any(x1))
	fmt.Println(format.Any(d))
	fmt.Println(format.Any([]int64{x1}))
	fmt.Println(format.Any([]time.Duration{d}))
	fmt.Println(reflect.ValueOf(d).Kind())

	type MyInt int
	var x2 MyInt = 7
	t2 := reflect.TypeOf(x2)
	fmt.Println("t2 type:", t2)
	fmt.Println("t2 kind:", t2.Kind())
	v2 := reflect.ValueOf(x2)
	fmt.Println("v2 value:", v2.Int())
	fmt.Println("v2 type:", v2.Type())
	fmt.Println("v2 kind:", v2.Kind())
	fmt.Println("v2.int type:", reflect.TypeOf(v2.Int()))

	var x3 float64 = 3.4
	p := reflect.ValueOf(&x3)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.Elem().CanSet())
}