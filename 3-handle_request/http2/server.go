package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"net/http"
)

type MyHandler struct {
}

func (myHandler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloWorld")
}

func main() {

	myHandler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &myHandler,
	}
	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServe()
}
