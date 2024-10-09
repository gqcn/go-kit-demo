package main

import (
	"log"
	"net/http"

	"go-kit-demo/gateway/internal/service"
	"go-kit-demo/gateway/internal/transport"
)

func main() {
	var (
		svc           = service.NewAddService()
		createHandler = transport.NewCreateHandler(svc)
		searchHandler = transport.NewSearchHandler(svc)
	)
	http.Handle("/create", createHandler)
	http.Handle("/search", searchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
