package main

import (
	"net/http"

	"github.com/mindriot101/goserver/server"
)

func main() {
	s := server.New()
	s.AddRoutes()
	http.ListenAndServe("127.0.0.1:8080", s)
}
