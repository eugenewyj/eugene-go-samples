package main

import (
	"fmt"
)

func main() {
	for pos, char := range "日本\x80语" {
		fmt.Printf("字符 %#U 从第 %d 个字节开始\n", char, pos)
	}
}
