package main

import (
	"net/http"
)

func main() {

	// 多路复用器进行重定向
	mux := http.NewServeMux()
	// 处理静态文件，去除/static/路径，并在public路径中查找文件
	files := http.FileServer(http.Dir("/2-chitchat/public"))
	mux.Handle("/static", http.StripPrefix("/static/", files))
	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)

	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
		//ReadTimeout: time.Duration(),
	}
	_ = server.ListenAndServe()
}
