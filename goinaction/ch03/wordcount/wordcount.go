package main

import (
	"fmt"
	"github.com/eugenewyj/go-sample/goinaction/ch03/words"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in you text. \n", count)
}
