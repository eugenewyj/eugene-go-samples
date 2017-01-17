package main

import (
	"net/http"
	"os"
	"log"
	"io"
)

func main() {
	r, err := http.Get(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	dest := io.MultiWriter(os.Stdout, file)
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
