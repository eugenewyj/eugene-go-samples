package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsDoneInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsDoneInfo)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		log.Fatal(err)
	}
}
