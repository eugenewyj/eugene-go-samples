package main

import (
	"log"
	"net/http"
	"github.com/eugenewyj/go-sample/goinaction/ch09/ch0914/handlers"
)

func main() {
	handlers.Routes()

	log.Println("Listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
