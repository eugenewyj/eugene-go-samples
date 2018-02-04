// 打印每个参数的索引及内容
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args[1:] {
		fmt.Printf("索引：%2d, 值：%s\n", index+1, value)
	}
}
