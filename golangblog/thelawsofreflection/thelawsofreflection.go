// golang反射法则 (https://blog.golang.org/laws-of-reflection)
// 1. Reflection goes from interface value to reflection object.
// 2. Refelction goes from reflection object to interface value.
// 3. To modify a reflection object, the value must be settable.
package thelawsofreflection

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type MyFloat64 float64

func main() {
	var x float64 = 3.4
	var mx1 MyFloat64 = MyFloat64(x)
	var x1 float64 = float64(mx1)
	fmt.Println(mx1)
	fmt.Println(x1)
	fmt.Println("x type :", reflect.TypeOf(x))
	fmt.Println("x type kind:", reflect.TypeOf(x).Kind())
	fmt.Println("x value :", reflect.ValueOf(x))

	var mx MyFloat64 = MyFloat64(x1)
	fmt.Println("mx type :", reflect.TypeOf(mx))
	fmt.Println("mx type kind:", reflect.TypeOf(mx).Kind())
	fmt.Println("mx value :", reflect.ValueOf(mx))
	fmt.Println("mx value type:", reflect.ValueOf(mx).Type())

	mxk := reflect.ValueOf(mx).Interface().(MyFloat64)
	fmt.Println("mxk:", mxk)

	var r io.Reader
	tty, _ := os.Open("/dev/tty")
	r = tty
	fmt.Println("r type:", reflect.TypeOf(r))
	fmt.Println("r type kind:", reflect.TypeOf(r).Kind())
	fmt.Println("tty type:", reflect.TypeOf(tty))
	fmt.Println("tty type kind:", reflect.TypeOf(tty).Kind())
}
