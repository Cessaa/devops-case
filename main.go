package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	host := os.Getenv("HOST")
	addr := fmt.Sprintf("%s:%s", host, port)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Write([]byte("Hi\n"))
		return
	})
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Write([]byte("pong\n"))
		return
	})
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("HTTP server started.\nAddr: \"%s\"", addr)
	go func(a string) {
		err := http.ListenAndServe(a, nil)
		panic(err)
	}(addr)
	<-s
}
