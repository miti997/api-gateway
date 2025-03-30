package main

import (
	"fmt"
	"net/http"

	"github.com/miti997/api-gateway/internal/routing"
)

func main() {
	sm := http.NewServeMux()

	r, _ := routing.NewRoute("GET", "/test", "https://jsonplaceholder.typicode.com/posts")
	sm.HandleFunc(r.GetPattern(), r.HandleRequest)

	r, _ = routing.NewRoute("GET", "/test/{id}", "https://jsonplaceholder.typicode.com/posts/{id}")
	sm.HandleFunc(r.GetPattern(), r.HandleRequest)

	s := &http.Server{
		Addr:    ":8081",
		Handler: sm,
	}

	fmt.Println("API Gateway is starting...")

	s.ListenAndServe()
}
