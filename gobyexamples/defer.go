package main

import (
	"fmt"
	"os"
)

// 参考：https://studygolang.com/articles/11907

func crateFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintf(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

// defer样例1：延迟调用函数
func example1() {
	defer func() {
		fmt.Println("样例1：后执行")
	}()
	fmt.Println("样例1：先执行")
}

// defer样例2： 释放已取得资源
func writeFileExample() {
	f := crateFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

// defer样例3：从故障中恢复
func panicRecoverExample() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("样例2：错误恢复")
		}
	}()
	fmt.Println("样例2：发生错误")
	panic("样例2：发生错误")
}

// defer样例4：延迟闭包
func delayClosure() {
	num := 42
	defer func() {
		fmt.Printf("样例4：num = %d\n", num)
	}()
	num = 13
}

// defer样例5：延迟闭包参数绑定
func delayClosureArgBind(i int) (n int) {
	defer func(j int) {
		fmt.Printf("样例5：j = %d\n", j)
		n = n + j
	}(i)

	i = i * 2
	fmt.Printf("样例5：i = %d\n", i)
	n = i

	return
}

// defer样例6： 延迟调用多个函数
func deployMultiFunc() {
	defer func() {
		fmt.Println("样例6：后调用")
	}()

	defer func() {
		fmt.Println("样例6：先调用")
	}()
}

type Car struct {
	model1 string
	model2 string
}

func (c Car) PrintModel1() {
	fmt.Println("样例7：非指针接受者model1=" + c.model1)
}

func (c *Car) PrintModel2() {
	fmt.Println("样例7：指针接受者model2=" + c.model2)
}

// 主函数
func main() {
	//writeFileExample()
	example1()
	panicRecoverExample()
	delayClosure()
	fmt.Printf("样例5：n=%d\n", delayClosureArgBind(10))
	deployMultiFunc()
	c := Car{model1: "model1:修改前", model2: "model2:修改前"}
	defer c.PrintModel2()
	defer c.PrintModel1()
	c.model1 = "model1:修改后"
	c.model2 = "model2:修改后"
}
