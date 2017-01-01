package main

import (
	"log"
	"os"

	_ "github.com/eugenewyj/go-sample/goinaction/ch02/matchers"
	"github.com/eugenewyj/go-sample/goinaction/ch02/search"
)

// init 方法在 main 方法之前调用。
func init()  {
	// 设置stdout为logging输出。
	log.SetOutput(os.Stdout)
}

// main 程序入口
func main() {
	// 执行search
 	search.Run("president")
}


