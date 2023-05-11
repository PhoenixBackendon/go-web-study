package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", getHeader)

	server := http.Server{
		Addr: "localhost:8080",
	}
	server.ListenAndServe()

}

func getHeader(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintln(w, k, "", v)
	}
}
