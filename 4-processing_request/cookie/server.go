package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{Addr: "localhost:8080"}
	http.HandleFunc("/setCookie1", setCookie1)
	http.HandleFunc("/setCookie2", setCookie2)
	http.HandleFunc("/getCookie1", getCookie1)
	http.HandleFunc("/getCookie2", getCookie2)
	server.ListenAndServe()
}

func setCookie1(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first",
		Value:    "gogogo",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second",
		Value:    "java!",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}

func setCookie2(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first",
		Value:    "gogogo",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second",
		Value:    "java!",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie1(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}

func getCookie2(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("second")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, cookie)
	fmt.Fprintln(w, cs)
}
